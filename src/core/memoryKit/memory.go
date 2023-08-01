package memoryKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/mem"
)

// GetUsedMemoryPercent
/*
@return e.g. (50.903940200805664, nil)
*/
func GetUsedMemoryPercent() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	return stat.UsedPercent, nil
}

// GetMemoryStat 获取当前瞬间的内存状态.
/*
PS:
(1) Total = Available + Used
*/
var GetMemoryStat func() (*mem.VirtualMemoryStat, error) = mem.VirtualMemory

// MemoryStatToString
/*
e.g.
() => "total: 32 GiB, available: 18 GiB, used: 14 GiB, free: 9.9 GiB, used percent: 43.06%", nil
*/
func MemoryStatToString(stat *mem.VirtualMemoryStat) string {
	return fmt.Sprintf("total: %s, available: %s, used: %s, free: %s, used percent: %.2f%%",
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		dataSizeKit.ToReadableStringWithIEC(stat.Available),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		stat.UsedPercent,
	)
}
