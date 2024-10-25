package utils

import (
	"encoding/json"
	"monutil/agent/config"
	"monutil/agent/logger"
	"monutil/agent/utils/cpu"
	"monutil/agent/utils/disk"
	hosts "monutil/agent/utils/host"
	loads "monutil/agent/utils/load"
	"monutil/agent/utils/mem"
	nets "monutil/agent/utils/net"
	"monutil/common/domain"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-netty/go-netty"
)

type Report interface {
	Connect(setupCodec func(channel netty.Channel))
	GetActive() bool
	SetStat(bool)
	Report(msg []byte)
}

func Monitors(cfg *config.Conf, cli Report) domain.MonStat {
	monstat := new(domain.MonStat)
	monstat.MsgType = "metrics"
	monstat.Msg.TimeStamp = time.Now()
	monstat.Msg.AgentId = cfg.Viper.GetString("agent.name")
	opts := GetMonitorOpts(cfg)
	for _, option := range opts {
		opt := option.(string)
		Monitor(opt, monstat)
	}
	b, err := json.Marshal(monstat)
	if err != nil {
		logger.Error("解析失败")
	}
	cli.Report(b)
	return *monstat
}

func Monitor(opt string, monstat *domain.MonStat) {
	if opt == "host" {
		monstat.Msg.Host = hosts.GetHost()
	} else if opt == "load" {
		monstat.Msg.Load = loads.GetLoad()
	} else if opt == "cpu" {
		monstat.Msg.CPU = cpu.GetCPU()
	} else if opt == "mem" {
		monstat.Msg.Mem = mem.GetMem()
	} else if opt == "disk" {
		monstat.Msg.Disk = disk.GetDisk()
	} else if opt == "net" {
		monstat.Msg.NetIO = nets.GetNet()
	}
}

func GetMonitorOpts(cfg *config.Conf) []interface{} {
	monset := mapset.NewSet()
	include := cfg.Viper.GetStringSlice("moniotr.include")
	if len(include) == 0 {
		monset.Add("host")
		monset.Add("load")
		monset.Add("cpu")
		monset.Add("mem")
		monset.Add("disk")
		monset.Add("net")
	} else {
		for _, opt := range include {
			monset.Add(opt)
		}
	}
	exclude := cfg.Viper.GetStringSlice("monitor.exclude")
	if len(exclude) > 0 {
		for _, opt := range exclude {
			monset.Remove(opt)
		}
	}
	return monset.ToSlice()
}
