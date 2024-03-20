package http

import (
	"encoding/json"
	"fmt"
	"github.com/cossim/hipush/api/http/v1/dto"
	"github.com/cossim/hipush/api/push"
	"github.com/cossim/hipush/pkg/consts"
	"github.com/cossim/hipush/pkg/notify"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) handleHuaweiPush(c *gin.Context, req *dto.PushRequest) error {
	service, err := h.factory.GetPushService(consts.Platform(req.Platform).String())
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return err
	}

	dataBytes, err := json.Marshal(req.Data)
	if err != nil {
		h.logger.Error(err, "Failed to marshal data")
		c.JSON(http.StatusBadRequest, Response{Code: http.StatusBadRequest, Msg: "invalid data", Data: nil})
		return err
	}
	var r dto.HuaweiPushRequestData
	if err := json.Unmarshal(dataBytes, &r); err != nil {
		h.logger.Error(err, "Failed to unmarshal data")
		c.JSON(http.StatusBadRequest, Response{Code: http.StatusBadRequest, Msg: "invalid data", Data: nil})
		return err
	}

	h.logger.Info("Handling push request", "platform", req.Platform, "appID", req.AppID, "tokens", req.Token, "req", r)

	fmt.Println("title => ", r.Title)

	rr := &notify.HMSPushNotification{
		AppID:       req.AppID,
		AppName:     req.AppName,
		Tokens:      req.Token,
		Development: r.Development,
		Title:       r.Title,
		Content:     r.Content,
		Category:    r.Category,
		Badge: notify.BadgeNotification{
			AddNum:     r.Badge.AddNum,
			SetNum:     r.Badge.SetNum,
			BadgeClass: r.Badge.Class,
		},
	}
	resp, err := service.Send(c, rr, &push.SendOptions{
		DryRun:        req.Option.DryRun,
		Retry:         req.Option.Retry,
		RetryInterval: req.Option.RetryInterval,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return err
	}

	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Msg: "Push notification send success", Data: resp})
	return nil
}
