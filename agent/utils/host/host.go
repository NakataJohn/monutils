package host

import "github.com/shirou/gopsutil/host"

func GetHost() *host.InfoStat {
	info, _ := host.Info()
	return info
}
