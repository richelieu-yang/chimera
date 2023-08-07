package memoryKit

import "github.com/shirou/gopsutil/v3/mem"

// GetMachineMemoryStats 获取当前瞬间的服务器内存状态.
/*
PS:
(1) Total = Available + Used
(2) UsedPercent: 内存使用率 e.g.50.903940200805664
*/
var GetMachineMemoryStats func() (*mem.VirtualMemoryStat, error) = mem.VirtualMemory
