/*
 * 心跳handler，
 * 当读超时时，向服务端发送心跳消息
 */

package handler

import (
	"encoding/json"
	"monutil/agent/logger"
	"monutil/common/domain"
	"strings"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/go-ping/ping"
)

const HeartBeat string = "Heartbeat"

var heartbeatMsg = &domain.Buf{
	MsgType: "heartbeat",
	Msg:     HeartBeat,
}

type HeartBeatHandler struct{}

func (h HeartBeatHandler) HandleEvent(ctx netty.EventContext, event netty.Event) {
	ip := strings.Split(ctx.Channel().RemoteAddr(), ":")[0]

	logger.Debug("Heartbeat HandleEvent", event)

	//只有当连接存活时出现写空闲则发送心跳
	//netty的IsActive方法无法准确判断网络异常,使用发送icmp包测试网络可达状态
	if _, _ok := event.(netty.WriteIdleEvent); _ok {
		pinger, _ := ping.NewPinger(ip)
		pinger.Count = 3
		pinger.Timeout = time.Second * 3
		pinger.SetPrivileged(true)
		err := pinger.Run()
		if err != nil {
			logger.Errorf("%s网络异常:%s", ip, err.Error())
		} else {
			logger.Infof(time.Now().Format("[2006-01-02 15:04:05] Send heartbeat_message to server"), HeartBeat)
			heartBeat, err := json.Marshal(heartbeatMsg)
			if err != nil {
				logger.Error(err)
			}
			ctx.Channel().Write(heartBeat)
		}
	}
}
