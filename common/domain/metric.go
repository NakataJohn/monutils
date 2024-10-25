package domain

import (
	"monutil/agent/utils/cpu"
	"monutil/agent/utils/disk"
	"monutil/agent/utils/mem"
	"time"

	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/net"
)

type MonStat struct {
	MsgType string `json:"msgType"`
	Msg     data   `json:"msg"`
}

type data struct {
	AgentId   string             `json:"agentId"`
	TimeStamp time.Time          `json:"timeStamp"`
	Host      *host.InfoStat     `json:"host,omitempty"`
	Load      *load.AvgStat      `json:"load,omitempty"`
	CPU       cpu.CPUInfo        `json:"cpu,omitempty"`
	Mem       mem.MemInfo        `json:"mem,omitempty"`
	Disk      []disk.DiskInfo    `json:"disk,omitempty"`
	NetIO     net.IOCountersStat `json:"netio,omitempty"`
}
