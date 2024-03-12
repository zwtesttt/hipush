package http

import (
	"encoding/json"
	"github.com/cossim/hipush/api/http/v1/dto"
	"github.com/cossim/hipush/pkg/consts"
	"github.com/cossim/hipush/pkg/notify"
	"github.com/cossim/hipush/pkg/push"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) handleVivoPush(c *gin.Context, req *dto.PushRequest) error {
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
	var r dto.VivoPushRequestData
	if err := json.Unmarshal(dataBytes, &r); err != nil {
		h.logger.Error(err, "Failed to unmarshal data")
		c.JSON(http.StatusBadRequest, Response{Code: http.StatusBadRequest, Msg: "invalid data", Data: nil})
		return err
	}

	h.logger.Info("Handling push request", "platform", req.Platform, "appID", req.AppID, "tokens", req.Token, "req", r)

	rr := &notify.VivoPushNotification{
		AppID:     req.AppID,
		RequestId: r.NotifyID,
		Tokens:    req.Token,
		Title:     r.Title,
		Message:   r.Content,
		Category:  r.Category,
		Data:      r.Data,
		ClickAction: &notify.VivoClickAction{
			Action:   r.ClickAction.Action,
			Url:      r.ClickAction.Url,
			Activity: r.ClickAction.Activity,
		},
		NotifyType:  r.NotifyType,
		TTL:         r.TTL,
		Retry:       0,
		SendOnline:  false,
		Foreground:  r.Foreground,
		Development: req.Option.Development,
	}
	if err := service.Send(c, rr, &push.SendOptions{
		DryRun:        req.Option.DryRun,
		Retry:         req.Option.Retry,
		RetryInterval: req.Option.RetryInterval,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Code: http.StatusBadRequest, Msg: err.Error(), Data: nil})
		return err
	}

	c.JSON(http.StatusOK, Response{Code: http.StatusOK, Msg: "Push notification send success", Data: nil})
	return nil
}
