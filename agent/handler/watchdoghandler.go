/*
 * 监听连接是否中断，中断连接进行重连
 *
 */

package handler

import (
	"fmt"
	"monutil/agent/logger"
	"monutil/agent/utils"
	"sync"
	"time"

	"github.com/go-netty/go-netty"
)

type WatchDogHandler struct {
	mu         sync.Mutex
	bootstrap  netty.Bootstrap
	Retry      time.Duration
	URI        string
	connecting bool
}

func NewWatchDogHandler(rtime time.Duration, url string) *WatchDogHandler {
	return &WatchDogHandler{
		Retry: rtime,
		URI:   url,
	}
}

func (w *WatchDogHandler) Watch(c utils.Report, bootstrap netty.Bootstrap) {

	w.bootstrap = bootstrap
	timer := time.NewTimer(w.Retry)

	for con_time := range timer.C {
		logger.Debug("重连时间：", con_time)
		w.connect(c)
		timer.Reset(w.Retry)
	}
}

func (w *WatchDogHandler) connect(c utils.Report) {
	w.mu.Lock()
	defer w.mu.Unlock()
	//正在连接则不重复连接
	if w.connecting {
		return
	}
	w.connecting = true
	ch, err := w.bootstrap.Connect(w.URI)
	w.connecting = false
	if err != nil {
		logger.Error(err)
		c.SetStat(false)
		return
	} else {
		c.SetStat(true)
	}

	select {
	case <-ch.Context().Done():
		c.SetStat(false)
		fmt.Println("ch.Context().Done()")
	case <-w.bootstrap.Context().Done():
		c.SetStat(false)
		fmt.Println("w.bootstrap.Context().Done()")
	}
}

/*
 * 断开连接
 */
func (w *WatchDogHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	logger.Error("inactive:", ctx.Channel().RemoteAddr(), ex)
	if ex.Error() == "EOF" {
		logger.Error("服务端已关闭")
		ctx.Channel().Close(ex)
	}

}

func (w *WatchDogHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	logger.Error("Exception:", ctx.Channel().RemoteAddr(), ex)
	if ex.Error() == "EOF" {
		logger.Error("服务端已关闭")
		ctx.Channel().Close(ex)
	}
}
