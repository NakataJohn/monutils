package report

import (
	"monutil/agent/config"
	"monutil/agent/handler"
	"monutil/agent/logger"
	"time"

	"github.com/go-netty/go-netty"
)

// client结构体
type Client struct {
	AID          string        //agentid
	Server       string        //server地址:端口
	Retrytimeout time.Duration //重连间隔
	Heartbeat    *heartbeat    //heartbeat setting
	IsActive     bool          //是否可用
}

type heartbeat struct {
	Timeout time.Duration //心跳超时时间:秒
	Retry   time.Duration //重试间隔:秒
}

// client的构造函数
func NewClient(cfg *config.Conf) *Client {
	return &Client{
		AID:    cfg.Viper.GetString("agent.name"),
		Server: cfg.Viper.GetString("server.host"),
		Heartbeat: &heartbeat{
			Timeout: cfg.Viper.GetDuration("heartbeat.timeout"),
			Retry:   cfg.Viper.GetDuration("heartbeat.retry"),
		},
		IsActive: false,
	}
}

func (c *Client) Connect(setupCodec func(channel netty.Channel)) {
	//重试最小为10s
	if c.Heartbeat.Retry < time.Duration(10*time.Second) {
		c.Heartbeat.Retry = time.Duration(10 * time.Second)
	}
	logger.Infof("断连重试间隔为：%v 秒", c.Heartbeat.Retry)
	//超时最小为30s
	if c.Heartbeat.Timeout < time.Duration(30*time.Second) {
		c.Heartbeat.Timeout = time.Duration(30 * time.Second)
	}
	wd := handler.NewWatchDogHandler(c.Heartbeat.Retry, c.Server)
	wd.Watch(c, netty.NewBootstrap(netty.WithClientInitializer(setupCodec)))
}

func (c Client) GetActive() bool {
	return c.IsActive
}

func (c *Client) SetStat(stat bool) {
	c.IsActive = stat
}

func (c Client) Report(msg []byte) {
	if c.GetActive() {
		handler.Ctx.Write(msg)
	} else {
		logger.Warn("服务连接已丢失")
	}
}
