package main

import (
	"fmt"
	"monutil/agent/config"
	"monutil/agent/handler"
	"monutil/agent/logger"
	"monutil/agent/report"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/codec/frame"
)

var (
	logo = `
 _____             _   _ _ _____             _   
|     |___ ___ _ _| |_|_| |  _  |___ ___ ___| |_ 
| | | | . |   | | |  _| | |     | . | -_|   |  _|
|_|_|_|___|_|_|___|_| |_|_|__|__|_  |___|_|_|_|  
                                |___|            
<Host_Monitor_Agent		::v1.0_beta>                                                        

`
	cfg = config.LoadConf
	// setup client pipeline initializer.

)

func main() {
	logger.Info("Agent is starting...")
	fmt.Println(logo)
	cli := report.NewClient(cfg)
	wd := handler.NewWatchDogHandler(cli.Heartbeat.Retry, cli.Server)
	setupCodec := func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(frame.DelimiterCodec(102400, "Mu$$", true)).
			AddLast(netty.WriteIdleHandler(time.Duration(30 * time.Second))).
			AddLast(handler.RegisterHandler{AID: cli.AID}).
			AddLast(handler.HeartBeatHandler{}).
			AddLast(wd)
	}
	report.TimeTask(cfg, cli)
	cli.Connect(setupCodec)

}
