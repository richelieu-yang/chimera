package memoryKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/mem"
)

type (
	MemoryStat mem.VirtualMemoryStat
)

func (stat MemoryStat) String() string {
	return fmt.Sprintf("total: %s, available: %s, used: %s, free: %s, used percent: %.2f%%",
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		dataSizeKit.ToReadableStringWithIEC(stat.Available),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		stat.UsedPercent,
	)
}

//// GetAvailableMemoryRatio 获取可用内存的比例
///*
//@return e.g. (0.4909677505493164, nil)
//
//“比例”和“比率”的区别？
//全班人数50人，男生30，女生20，那男生的比例就是30/50，同理女生的就是20/50，那么男女的比率是什么呢，是30/20。
//*/
//func GetAvailableMemoryRatio() (float64, error) {
//	stat, err := mem.VirtualMemory()
//	if err != nil {
//		return 0.0, err
//	}
//	tmp := floatKit.Div(float64(stat.Available), float64(stat.Total))
//	return tmp, nil
//}
//
//// GetAvailableMemoryPercent
///*
//@return e.g. (49.09634590148926, nil)
//*/
//func GetAvailableMemoryPercent() (float64, error) {
//	stat, err := mem.VirtualMemory()
//	if err != nil {
//		return 0.0, err
//	}
//	return 100 - stat.UsedPercent, nil
//}
//
//// GetUsedMemoryRatio
///*
//@return e.g. (0.509037971496582, nil)
//*/
//func GetUsedMemoryRatio() (float64, error) {
//	stat, err := mem.VirtualMemory()
//	if err != nil {
//		return 0.0, err
//	}
//	tmp := floatKit.Div(float64(stat.Used), float64(stat.Total))
//	return tmp, nil
//}

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

// GetMemoryStat 获取当前瞬间的内存状态
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
