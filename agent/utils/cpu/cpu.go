package cpu

import (
	"github.com/shirou/gopsutil/cpu"
)

type CPUInfo struct {
	InfoStat `json:"infoStat"`
	Percent  float64       `json:"percent"`
	TimeStat cpu.TimesStat `json:"timeStat"`
}

type InfoStat struct {
	ModelName string `json:"modelName"`
	Cores     int    `json:"logicalCores"`
}

func GetCPU() CPUInfo {
	cpustat, _ := cpu.Info()
	cores, _ := cpu.Counts(true)
	f, _ := cpu.Percent(0, false)
	timestat, _ := cpu.Times(false)

	cpuinfo := CPUInfo{
		InfoStat: InfoStat{
			ModelName: cpustat[0].ModelName,
			Cores:     cores,
		},
		Percent:  f[0],
		TimeStat: timestat[0],
	}

	return cpuinfo
}
