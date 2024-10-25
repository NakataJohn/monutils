package mem

import "github.com/shirou/gopsutil/mem"

type MemInfo struct {
	VirtualMemory *mem.VirtualMemoryStat `json:"virtualMemory"`
	SwapMemory    *mem.SwapMemoryStat    `json:"swapMemory"`
}

func GetMem() MemInfo {
	vmemory, _ := mem.VirtualMemory()
	smemory, _ := mem.SwapMemory()

	return MemInfo{
		VirtualMemory: vmemory,
		SwapMemory:    smemory,
	}
}
