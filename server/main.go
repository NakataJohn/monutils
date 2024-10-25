package main

import (
	"fmt"
	"monutil/server/config"
	"monutil/server/handler"

	"time"

	"monutil/server/logger"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/format"
	"github.com/go-netty/go-netty/codec/frame"
)

var (
	logo = `                                            
 _____             _   _ _ _____                     
|     |___ ___ _ _| |_|_| |   __|___ ___ _ _ ___ ___ 
| | | | . |   | | |  _| | |__   | -_|  _| | | -_|  _|
|_|_|_|___|_|_|___|_| |_|_|_____|___|_|  \_/|___|_|                                                      
<Host_Monitor_Server		::v1.0_beta>                                                        

`
	mgr       = handler.CtxManager
	cfg       = config.LoadConf
	hbtimeout = cfg.Viper.GetInt64("server.hbtimeout")
	handlers  = []netty.Handler{
		// frame.PacketCodec(1024),
		frame.DelimiterCodec(102400, "Mu$$", true),
		format.JSONCodec(true, false),
		netty.ReadIdleHandler(time.Duration(hbtimeout * int64(time.Second))),
		mgr,
		handler.HbHandle{},
		handler.DoHandler{},
	}
)

func start(handlers []netty.Handler) {
	// read listen addr from config file
	listen := cfg.Viper.GetString("server.listen")
	childInitializer := func(channel netty.Channel) {
		pipeline := channel.Pipeline()
		for _, handler := range handlers {
			pipeline.AddLast(handler)
		}
	}
	// create Bootstrap & port listen & accpect connection from agent
	netty.NewBootstrap(netty.WithChildInitializer(childInitializer)).
		Listen(listen).Sync()
	logger.Infof("服务监听：%s", listen)
}

func main() {
	logger.Info("Host Monitor Server is starting...")
	fmt.Println(logo)
	start(handlers)
}
