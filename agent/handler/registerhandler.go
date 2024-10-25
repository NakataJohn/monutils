/*
 * 注册handler，
 * 向服务端发送注册消息
 */

package handler

import (
	"encoding/json"
	"monutil/agent/logger"
	"monutil/common/domain"

	"github.com/go-netty/go-netty"
)

var (
	Ctx     netty.HandlerContext
	initMsg = &domain.Buf{
		MsgType: "regist",
	}
)

type RegisterHandler struct {
	AID string //agent ID
}

func (r RegisterHandler) HandleActive(ctx netty.ActiveContext) {
	logger.Infof("连接服务%s成功", ctx.Channel().RemoteAddr())
	initMsg.Msg = r.AID
	b, err := json.Marshal(initMsg)
	if err != nil {
		logger.Error("解析失败")
	}
	ctx.Write(b)
	Ctx = ctx
	ctx.HandleActive()
}

func (r RegisterHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	logger.Infof("客户端收到消息：%v", message)
}
