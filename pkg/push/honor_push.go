package push

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cossim/hipush/config"
	hClient "github.com/cossim/hipush/pkg/client/push"
	"github.com/cossim/hipush/pkg/consts"
	"github.com/cossim/hipush/pkg/notify"
	"github.com/cossim/hipush/pkg/status"
	"github.com/go-logr/logr"
	"log"
	"net/http"
)

var (
	// MaxConcurrentHonorPushes pool to limit the number of concurrent iOS pushes
	MaxConcurrentHonorPushes = make(chan struct{}, 100)
)

// HonorService 荣耀推送，实现了 PushService 接口
type HonorService struct {
	clients map[string]*hClient.HonorPushClient
	status  *status.StateStorage
	logger  logr.Logger
}

func NewHonorService(cfg *config.Config, logger logr.Logger) *HonorService {
	s := &HonorService{
		clients: make(map[string]*hClient.HonorPushClient),
		status:  status.StatStorage,
		logger:  logger,
	}

	for _, v := range cfg.Honor {
		if !v.Enabled {
			continue
		}
		if v.Enabled && (v.AppID == "" || v.ClientID == "" || v.ClientSecret == "") {
			panic("push not enabled or misconfigured")
		}
		s.clients[v.AppID] = hClient.NewHonorPush(v.ClientID, v.ClientSecret)
	}

	return s
}

func (h *HonorService) Send(ctx context.Context, request interface{}, opt ...SendOption) error {
	req, ok := request.(*notify.HonorPushNotification)
	if !ok {
		return errors.New("invalid request")
	}

	so := &SendOptions{}
	so.ApplyOptions(opt)

	if err := h.checkNotification(req); err != nil {
		return err
	}

	notification := h.buildAndroidNotification(req)

	if so.DryRun {
		return nil
	}

	send := func(ctx context.Context, token string) (*Response, error) {
		return h.send(ctx, req.AppID, token, notification)
	}
	return RetrySend(ctx, send, req.Tokens, so.Retry, so.RetryInterval, 100)

	//for {
	//	newTokens, err := h.send(ctx, req.AppID, req.Tokens, notification)
	//	if err != nil {
	//		log.Printf("send error => %v", err)
	//		es = append(es, err)
	//	}
	//	// 如果有重试的 Token，并且未达到最大重试次数，则进行重试
	//	if len(newTokens) > 0 && retryCount < maxRetry {
	//		retryCount++
	//		req.Tokens = newTokens
	//	} else {
	//		break
	//	}
	//}
	//
	//var errorMsgs []string
	//for _, err := range es {
	//	errorMsgs = append(errorMsgs, err.Error())
	//}
	//if len(errorMsgs) > 0 {
	//	return fmt.Errorf("%s", strings.Join(errorMsgs, ", "))
	//}
	//return nil
}

func (h *HonorService) send(ctx context.Context, appid string, token string, notification *hClient.SendMessageRequest) (*Response, error) {
	client, ok := h.clients[appid]
	if !ok {
		return nil, errors.New("invalid appid or appid push is not enabled")
	}

	h.status.AddHonorTotal(1)

	resp := &Response{Code: Fail}
	notification.Token = []string{token}
	res, err := client.SendMessage(ctx, appid, notification)
	if err != nil {
		log.Printf("honor send error: %s", err)
		h.status.AddHonorFailed(1)
		resp.Msg = err.Error()
	} else if res != nil && res.Code != http.StatusOK {
		if len(res.Data.ExpireTokens) > 0 {
			log.Printf("honor send expire tokens: %s", res.Data.ExpireTokens)
		}
		log.Printf("honor send error: %s", res.Message)
		h.status.AddHonorFailed(1)
		err = errors.New(res.Message)
		resp.Code = res.Code
		resp.Msg = res.Message
	} else {
		log.Printf("honor send success: %s", res.Message)
		h.status.AddHonorSuccess(1)
		resp.Code = Success
		resp.Msg = res.Message
		resp.Data = res.Data
	}

	return resp, err

	//var es []error
	//
	//for _, token := range tokens {
	//	// occupy push slot
	//	MaxConcurrentHuaweiPushes <- struct{}{}
	//	wg.Add(1)
	//	h.status.AddHonorTotal(1)
	//	go func(notification *hClient.SendMessageRequest, token string) {
	//		defer func() {
	//			// free push slot
	//			<-MaxConcurrentHuaweiPushes
	//			wg.Done()
	//		}()
	//		res, err := client.SendMessage(ctx, appid, notification)
	//		if err != nil || (res != nil && res.Code != 200) {
	//			if err == nil {
	//				err = errors.New(res.Content)
	//			} else {
	//				es = append(es, err)
	//			}
	//			// 记录失败的 Token
	//			if res != nil && res.Code != 200 {
	//				newTokens = append(newTokens, res.Data.FailTokens...)
	//			}
	//			if len(res.Data.ExpireTokens) > 0 {
	//				log.Printf("honor send expire tokens: %s", res.Data.ExpireTokens)
	//			}
	//			log.Printf("honor send error: %s", err)
	//			h.status.AddHonorFailed(1)
	//
	//		} else {
	//			log.Printf("honor send success: %s", res.Content)
	//			h.status.AddHonorSuccess(1)
	//		}
	//	}(notification, token)
	//}
	//wg.Wait()
	//if len(es) > 0 {
	//	var errorStrings []string
	//	for _, err := range es {
	//		errorStrings = append(errorStrings, err.Error())
	//	}
	//	allErrorsString := strings.Join(errorStrings, ", ")
	//	return nil, errors.New(allErrorsString)
	//}
	//return newTokens, nil
}

func (h *HonorService) checkNotification(req *notify.HonorPushNotification) error {
	if len(req.Tokens) == 0 {
		return errors.New("tokens cannot be empty")
	}

	if req.Title == "" {
		return errors.New("title cannot be empty")
	}

	if req.Content == "" {
		return errors.New("content cannot be empty")
	}

	return nil
}

func (h *HonorService) buildAndroidNotification(req *notify.HonorPushNotification) *hClient.SendMessageRequest {
	// 构建通知栏消息
	notification := &hClient.Notification{
		Title: req.Title,
		Body:  req.Content,
		Image: req.Image,
	}

	// 构建 Android 平台的通知消息
	androidNotification := &hClient.AndroidNotification{
		Title: req.Title,
		Body:  req.Content,
		Image: req.Image,
		Badge: &hClient.BadgeNotification{
			AddNum:     req.Badge.AddNum,
			SetNum:     req.Badge.SetNum,
			BadgeClass: req.Badge.BadgeClass,
		},
		ClickAction: &hClient.ClickAction{
			Type:   req.ClickAction.Action,
			Intent: req.ClickAction.Activity,
			URL:    req.ClickAction.Url,
			Action: req.ClickAction.Activity,
		},
	}

	var targetUserType int

	if req.Development {
		targetUserType = 1
	}

	// 将 map 转换为 JSON 字符串
	data, err := json.Marshal(req.Data)
	if err != nil {
		log.Printf("Error marshalling data: %v", err)
	}

	// 构建 Android 平台消息推送配置
	androidConfig := &hClient.AndroidConfig{
		TTL:            req.TTL,      // 设置消息缓存时间
		BiTag:          "",           // 设置批量任务消息标识
		Data:           string(data), // 设置自定义消息负载
		Notification:   androidNotification,
		TargetUserType: targetUserType, // 设置目标用户类型
	}

	// 构建发送消息请求
	sendMessageReq := &hClient.SendMessageRequest{
		Data:         string(data), // 设置自定义消息负载
		Notification: notification,
		Android:      androidConfig,
		Token:        req.Tokens,
	}

	return sendMessageReq
}

func (h *HonorService) Name() string {
	return consts.PlatformHonor.String()
}