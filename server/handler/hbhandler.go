package handler

import (
	"errors"
	"fmt"
	"monutil/server/logger"

	"github.com/go-netty/go-netty"
)

type HbHandle struct{}

func (HbHandle) HandleRead(ctx netty.InboundContext, message netty.Message) {
	msg := message.(map[string]interface{})
	// bufMsg := []byte(message.(string))

	if msg["msgType"] == "Heartbeat" {
		logger.Info(msg["msg"], "From", ctx.Channel().RemoteAddr())
	} else {
		ctx.HandleRead(message)
	}
}

func (HbHandle) HandleEvent(ctx netty.EventContext, evt netty.Event) {
	if _, ok := evt.(netty.ReadIdleEvent); ok {
		err := errors.New("heartbeat was losted")
		errmsg := fmt.Sprintf("Heartbeat of %s was losted", ctx.Channel().RemoteAddr())
		logger.Error(errmsg)
		ctx.Channel().Close(err)
	}
}
