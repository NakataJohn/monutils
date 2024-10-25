package load

import "github.com/shirou/gopsutil/load"

func GetLoad() *load.AvgStat {
	load, _ := load.Avg()
	return load
}
