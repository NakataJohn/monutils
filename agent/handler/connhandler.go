package handler

import (
	"sync"

	"github.com/go-netty/go-netty"
)

type ConnHandler struct {
	Ctx    netty.HandlerContext
	Getctx func(netty.HandlerContext)
	_mutex sync.RWMutex
}

func (c *ConnHandler) HandleActive(ctx netty.ActiveContext) {
	c._mutex.RLock()
	c.Getctx(ctx)
	c._mutex.RUnlock()
	ctx.HandleActive()
}
