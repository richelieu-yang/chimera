package runtimeKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/floatKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/disk"
)

type (
	// DiskStat
	/*
		Total = Used + Free
	*/
	DiskStat disk.UsageStat
)

// GetFreePercent
/*
@return e.g. 65.18472380640327
*/
func (stat *DiskStat) GetFreePercent() float64 {
	return floatKit.Mul(100, floatKit.Div(float64(stat.Free), float64(stat.Total)))
}

// String
/*
@return e.g. "free: 213 GiB, used: 247 GiB, total: 460 GiB, free percent: 46.27%"
*/
func (stat *DiskStat) String() string {
	return fmt.Sprintf("free: %s, used: %s, total: %s, free percent: %.2f%%",
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		stat.GetFreePercent(),
	)
}

// GetDiskStat
/*
参考: golang 获取cpu 内存 硬盘 使用率 信息 进程信息 https://blog.csdn.net/whatday/article/details/109620192
*/
func GetDiskStat() (*DiskStat, error) {
	parts, err := disk.Partitions(true)
	if err != nil {
		return nil, err
	}

	for _, part := range parts {
		if part.Mountpoint != "/" {
			continue
		}
		stat, err := disk.Usage(part.Mountpoint)
		if err != nil {
			return nil, err
		}
		return (*DiskStat)(stat), nil
	}
	return nil, errorKit.Newf("fail to get disk stat with parts(length: %d)", len(parts))
}
