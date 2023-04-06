package runtimeKit

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/mem"
)

type (
	// MemoryStat Total = Available + Used
	MemoryStat mem.VirtualMemoryStat
)

// GetAvailablePercent
/*
e.g. 51.035356521606445
*/
func (stat MemoryStat) GetAvailablePercent() float64 {
	return floatKit.Mul(100, floatKit.Div(float64(stat.Available), float64(stat.Total)))
}

func (stat MemoryStat) String() string {
	return fmt.Sprintf("available: %s, free: %s, used: %s, total: %s, available percent: %.2f%%",
		dataSizeKit.ToReadableStringWithIEC(stat.Available),
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		stat.GetAvailablePercent(),
	)
}

// GetAvailableMemoryRatio 获取可用内存的比例
/*
@return e.g. (0.4909677505493164, nil)

“比例”和“比率”的区别？
全班人数50人，男生30，女生20，那男生的比例就是30/50，同理女生的就是20/50，那么男女的比率是什么呢，是30/20。
*/
func GetAvailableMemoryRatio() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	tmp := floatKit.Div(float64(stat.Available), float64(stat.Total))
	return tmp, nil
}

// GetAvailableMemoryPercent
/*
@return e.g. (49.09634590148926, nil)
*/
func GetAvailableMemoryPercent() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	return 100 - stat.UsedPercent, nil
}

// GetUsedMemoryRatio
/*
@return e.g. (0.509037971496582, nil)
*/
func GetUsedMemoryRatio() (float64, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0.0, err
	}
	tmp := floatKit.Div(float64(stat.Used), float64(stat.Total))
	return tmp, nil
}

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
func GetMemoryStat() (*MemoryStat, error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return (*MemoryStat)(stat), nil
}
