package push

import (
	"context"
	"errors"
	"fmt"
	"github.com/cossim/hipush/config"
	"github.com/cossim/hipush/internal/notify"
	vp "github.com/cossim/vivo-push"
	"log"
	"strings"
	"sync"
)

// VivoService 实现vivo推送，必须实现PushService接口
type VivoService struct {
	clients map[string]*vp.VivoPush
}

func NewVivoService(cfg *config.Config) (*VivoService, error) {
	s := &VivoService{
		clients: map[string]*vp.VivoPush{},
	}

	for _, v := range cfg.Vivo {
		if !v.Enabled || v.Enabled && (v.AppID == "" || v.AppKey == "" || v.AppSecret == "") {
			return nil, errors.New("push not enabled or misconfigured")
		}
		client, err := vp.NewClient(v.AppID, v.AppKey, v.AppSecret)
		if err != nil {
			return nil, err
		}
		s.clients[v.AppID] = client
	}

	return s, nil
}

func (v *VivoService) Send(ctx context.Context, request interface{}) error {
	req, ok := request.(*notify.VivoPushNotification)
	if !ok {
		return errors.New("invalid request")
	}

	var (
		retry      = req.Retry
		maxRetry   = retry
		retryCount = 0
	)

	// 重试计数
	if maxRetry <= 0 {
		maxRetry = DefaultMaxRetry // 设置一个默认的最大重试次数
	}
	if retry > 0 && retry < maxRetry {
		maxRetry = retry
	}

	var es []error

	for {
		newTokens, err := v.send(req)
		if err != nil {
			log.Printf("sendNotifications error => %v", err)
			es = append(es, err)
		}
		// 如果有重试的 Token，并且未达到最大重试次数，则进行重试
		if len(newTokens) > 0 && retryCount < maxRetry {
			retryCount++
			req.Tokens = newTokens
		} else {
			break
		}
	}

	var errorMsgs []string
	for _, err := range es {
		errorMsgs = append(errorMsgs, err.Error())
	}
	if len(errorMsgs) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMsgs, ", "))
	}
	return nil
}

func (v *VivoService) send(req *notify.VivoPushNotification) ([]string, error) {
	var newTokens []string
	notification, err := v.buildNotification(req)
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup

	client, ok := v.clients[req.AppID]
	if !ok {
		return nil, errors.New("invalid appid or appid push is not enabled")
	}

	var es []error

	for _, token := range req.Tokens {
		// occupy push slot
		MaxConcurrentIOSPushes <- struct{}{}
		wg.Add(1)
		go func(notification *vp.Message, token string) {
			defer func() {
				// free push slot
				<-MaxConcurrentIOSPushes
				wg.Done()
			}()

			notification.RegId = token
			res, err := client.Send(notification, token)
			if err != nil || (res != nil && res.Result != 0) {
				if err == nil {
					err = errors.New(res.Desc)
				} else {
					es = append(es, err)
				}
				// 记录失败的 Token
				if res != nil && res.Result != 0 {
					newTokens = append(newTokens, token)
				}
				log.Printf("vivo send error: %s", err)
			} else {
				log.Printf("vivo send success: %s", res.Desc)
			}
		}(notification, token)
	}
	wg.Wait()
	if len(es) > 0 {
		var errorStrings []string
		for _, err := range es {
			errorStrings = append(errorStrings, err.Error())
		}
		allErrorsString := strings.Join(errorStrings, ", ")
		return nil, errors.New(allErrorsString)
	}
	return newTokens, nil
}

func (v *VivoService) buildNotification(req *notify.VivoPushNotification) (*vp.Message, error) {
	// 检查 tokens 是否为空
	if len(req.Tokens) == 0 {
		return nil, errors.New("tokens cannot be empty")
	}

	if req.Title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if req.Message == "" {
		return nil, errors.New("message cannot be empty")
	}

	// 设置默认的 ClickAction
	defaultClickAction := &notify.VivoClickAction{
		Action:  1,
		Content: "",
	}

	// 检查 ClickAction 是否为空，为空则使用默认值
	clickAction := req.ClickAction
	if clickAction == nil {
		clickAction = defaultClickAction
	}

	// 检查 NotifyType 是否为有效值
	if req.NotifyType != 0 && req.NotifyType < 1 || req.NotifyType > 4 {
		return nil, errors.New("invalid notify type")
	}

	if req.NotifyType == 0 {
		req.NotifyType = 2
	}

	var pushMode int
	if req.Development {
		pushMode = 1
	}

	if req.TTL == 0 {
		req.TTL = 60
	}

	message := &vp.Message{
		RegId:           strings.Join(req.Tokens, ","),
		NotifyType:      req.NotifyType,
		Title:           req.Title,
		Content:         req.Message,
		TimeToLive:      int64(req.TTL),
		SkipType:        clickAction.Action,
		SkipContent:     clickAction.Content,
		NetworkType:     -1,
		ClientCustomMap: req.Data,
		//Extra:           req.Data.ExtraMap(),
		RequestId:      req.RequestId,
		PushMode:       pushMode, // 默认为正式推送
		ForegroundShow: req.Foreground,
	}
	return message, nil
}

func (v *VivoService) MulticastSend(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) Subscribe(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) Unsubscribe(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) SendToTopic(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) SendToCondition(ctx context.Context, req interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) CheckDevice(ctx context.Context, req interface{}) bool {
	//TODO implement me
	panic("implement me")
}

func (v *VivoService) Name() string {
	//TODO implement me
	panic("implement me")
}