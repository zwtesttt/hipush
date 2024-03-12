package push

import (
	"context"
	"encoding/json"
	"errors"
	mzp "github.com/cossim/go-meizu-push-sdk"
	"github.com/cossim/hipush/config"
	"github.com/cossim/hipush/pkg/consts"
	"github.com/cossim/hipush/pkg/notify"
	"github.com/cossim/hipush/pkg/status"
	"github.com/go-logr/logr"
	"log"
)

var (
	MaxConcurrentMeizuPushes = make(chan struct{}, 100)
)

// MeizuService 实现魅族推送，实现 PushService 接口
type MeizuService struct {
	clients map[string]func(token, message string) mzp.PushResponse
	status  *status.StateStorage
	logger  logr.Logger
}

func NewMeizuService(cfg *config.Config, logger logr.Logger) *MeizuService {
	s := &MeizuService{
		clients: make(map[string]func(token, message string) mzp.PushResponse),
		status:  status.StatStorage,
		logger:  logger,
	}

	for _, v := range cfg.Meizu {
		if !v.Enabled || v.Enabled && (v.AppID == "" || v.AppKey == "") {
			panic("push not enabled or misconfigured")
		}
		s.clients[v.AppID] = func(token, message string) mzp.PushResponse {
			appid := v.AppID
			appkey := v.AppKey
			return mzp.PushNotificationMessageByPushId(appid, token, message, appkey)
		}
	}

	return s
}

func (m *MeizuService) Send(ctx context.Context, request interface{}, opt ...SendOption) error {
	req, ok := request.(*notify.MeizuPushNotification)
	if !ok {
		return errors.New("invalid request")
	}

	so := &SendOptions{}
	so.ApplyOptions(opt)

	if err := m.checkNotification(req); err != nil {
		return err
	}

	notification, err := m.buildNotification(req)
	if err != nil {
		return err
	}

	if so.DryRun {
		return nil
	}

	send := func(ctx context.Context, token string) (*Response, error) {
		return m.send(req.AppID, token, notification)
	}
	return RetrySend(ctx, send, req.Tokens, so.Retry, so.RetryInterval, 100)

	//for {
	//	newTokens, err := m.send(req.AppID, req.Tokens, notification)
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

func (m *MeizuService) send(appid string, token string, message string) (*Response, error) {
	pushFunc, ok := m.clients[appid]
	if !ok {
		return nil, errors.New("invalid appid or appid push is not enabled")
	}

	m.status.AddMeizuTotal(1)

	var err error
	resp := &Response{}
	res := pushFunc(token, message)
	if res.GetCode() != Success {
		log.Printf("meizu send error: %s", res.GetMessage())
		m.status.AddMeizuFailed(1)
		err = errors.New(res.GetMessage())
		resp.Code = Fail
		resp.Msg = res.GetMessage()
	} else {
		log.Printf("meizu send success code: %v msg: %s", res.GetCode(), res.GetMessage())
		m.status.AddMeizuSuccess(1)
		resp.Code = Success
		resp.Msg = res.GetMessage()
	}

	return resp, err

	//var es []error
	//
	//for _, token := range tokens {
	//	// occupy push slot
	//	MaxConcurrentXiaomiPushes <- struct{}{}
	//	wg.Add(1)
	//	m.status.AddMeizuTotal(1)
	//	go func(notification string, token string) {
	//		defer func() {
	//			// free push slot
	//			<-MaxConcurrentXiaomiPushes
	//			wg.Done()
	//		}()
	//
	//		fmt.Println("notification => ", notification)
	//		res := pushFunc(token, notification)
	//		if res.GetCode() != 200 {
	//			es = append(es, errors.New(res.GetMessage()))
	//			log.Printf("oppo send error: %s", res.GetMessage())
	//			// 记录失败的 Token
	//			newTokens = append(newTokens, token)
	//			m.status.AddMeizuFailed(1)
	//		} else {
	//			log.Printf("oppo send success code: %v msg: %s", res.GetCode(), res.GetMessage())
	//			m.status.AddMeizuSuccess(1)
	//		}
	//	}(message, token)
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

func (m *MeizuService) checkNotification(req *notify.MeizuPushNotification) error {
	if len(req.Tokens) == 0 {
		return errors.New("tokens cannot be empty")
	}

	if req.Title == "" {
		return errors.New("title cannot be empty")
	}

	if req.Content == "" {
		return errors.New("message cannot be empty")
	}

	if req.IsScheduled && (req.ScheduledStartTime == "" || req.ScheduledEndTime == "") {
		return errors.New("scheduled time cannot be empty")
	}

	return nil
}

func (m *MeizuService) buildNotification(req *notify.MeizuPushNotification) (string, error) {
	msg := mzp.BuildNotificationMessage()
	msg.NoticeBarInfo.Title = req.Title
	msg.NoticeBarInfo.Content = req.Content
	msg.ClickTypeInfo = mzp.ClickTypeInfo{
		ClickType:  req.ClickAction.Action,
		Url:        req.ClickAction.Url,
		Parameters: req.ClickAction.Parameters,
		Activity:   req.ClickAction.Activity,
	}

	offLine := 0
	if req.OffLine {
		offLine = 1
	}
	msg.PushTimeInfo = mzp.PushTimeInfo{
		OffLine:   offLine,
		ValidTime: req.TTL,
	}

	message, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func (m *MeizuService) Name() string {
	return consts.PlatformMeizu.String()
}