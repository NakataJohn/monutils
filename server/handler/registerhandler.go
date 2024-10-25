package handler

import (
	"errors"
	"fmt"
	"io"
	"monutil/server/logger"
	"sync"

	"github.com/go-netty/go-netty"
)

var CtxManager Manager

func init() {
	CtxManager = NewManager()
}

type Manager interface {
	netty.ActiveHandler
	netty.InboundHandler
	netty.InactiveHandler
	Size() int
	Context(aid string) (netty.HandlerContext, bool)
}

func NewManager() Manager {
	return &sessionManager{
		_sessions: make(map[string]netty.HandlerContext, 64),
	}
}

type sessionManager struct {
	_sessions map[string]netty.HandlerContext
	_mutex    sync.RWMutex
}

func (s *sessionManager) Size() int {
	s._mutex.RLock()
	size := len(s._sessions)
	s._mutex.RUnlock()
	return size
}

func (s *sessionManager) Context(aid string) (netty.HandlerContext, bool) {
	s._mutex.RLock()
	ctx, ok := s._sessions[aid]
	s._mutex.RUnlock()
	return ctx, ok
}

func (s *sessionManager) HandleActive(ctx netty.ActiveContext) {
	logger.Infof("Agent %v is connected!", ctx.Channel().RemoteAddr())
	ctx.HandleActive()
}

// 添加session
func (s *sessionManager) HandleRead(ctx netty.InboundContext, message netty.Message) {
	msg := message.(map[string]interface{})

	if msg["msgType"] == "regist" {
		sid := msg["msg"].(string)
		if _, ok := s._sessions[sid]; !ok {
			s._mutex.Lock()
			s._sessions[sid] = ctx
			logger.Infof("%v 注册成功", sid)
			// gomysql.DB.Model(&gomysql.Agent{}).Where("name=?", sid).Update("status", "E")
			s._mutex.Unlock()
		} else {
			msg := fmt.Sprintf("%v 已存在请勿重复启动相同agentname的agent，或者请先关闭已有的agent后再启动", sid)
			logger.Error(msg)
			ctx.Write(msg)
			ctx.Close(errors.New("重复注册"))
		}
	} else {
		ctx.HandleRead(message)
	}
}

func (*sessionManager) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	if ex.Error() == io.EOF.Error() {
		ctx.Close(ex)
	} else {
		logger.Errorf("程序异常:%s", ex.Error())
		ctx.Close(ex)
	}
}

// 删除session
func (s *sessionManager) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	for sid, dctx := range s._sessions {
		if ctx.Channel().ID() == dctx.Channel().ID() {
			s._mutex.Lock()
			delete(s._sessions, sid)
			logger.Infof(fmt.Sprintf("%v 已注销", sid))
			// gomysql.DB.Model(&gomysql.Agent{}).Where("name=?", sid).Update("status", "D")
			s._mutex.Unlock()
			ctx.HandleInactive(ex)
		}
	}
}
