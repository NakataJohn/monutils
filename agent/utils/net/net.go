package net

import "github.com/shirou/gopsutil/net"

func GetNet() net.IOCountersStat {
	ioStat, _ := net.IOCounters(false)
	return ioStat[0]
}
