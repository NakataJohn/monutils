package disk

import (
	"github.com/shirou/gopsutil/disk"
)

type DiskInfo struct {
	Mountpoint string          `json:"mountpoint"`
	Fstype     string          `json:"fstype"`
	UsageStat  *disk.UsageStat `json:"usageStat"`
}

func GetDisk() (disksInfo []DiskInfo) {
	devs, _ := disk.Partitions(true)
	for _, d := range devs {
		mp := d.Mountpoint
		fs := d.Fstype
		// dev := d.Device
		useage, _ := disk.Usage(mp)
		diskinfo := DiskInfo{
			Mountpoint: mp,
			Fstype:     fs,
			UsageStat:  useage,
		}
		disksInfo = append(disksInfo, diskinfo)
	}
	return disksInfo
}
