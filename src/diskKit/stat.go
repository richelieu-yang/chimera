package diskKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/floatKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/disk"
	"runtime"
)

type (
	// DiskStat
	/*
		Total = Used + Free
	*/
	DiskStat struct {
		*disk.UsageStat

		Path string `json:"path"`
	}
)

// GetUsedPercent
/*
@return e.g. 53.7314129274183
*/
func (stat *DiskStat) GetUsedPercent() float64 {
	return floatKit.Mul(100, floatKit.Div(float64(stat.Used), float64(stat.Total)))
}

// GetFreePercent
/*
@return e.g. 46.2685870725817
*/
func (stat *DiskStat) GetFreePercent() float64 {
	return floatKit.Mul(100, floatKit.Div(float64(stat.Free), float64(stat.Total)))
}

// String
/*
@return e.g. "path: /, free: 213 GiB, used: 247 GiB, total: 460 GiB, free percent: 46.27%"
*/
func (stat *DiskStat) String() string {
	return fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, free percent: %.2f%%",
		stat.Path,
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		stat.GetFreePercent(),
	)
}

// GetDiskStat
/*
PS:
(1) Mac（Linux），查看磁盘空间的命令: df -h

golang 获取cpu 内存 硬盘 使用率 信息 进程信息
	https://blog.csdn.net/whatday/article/details/109620192
*/
func GetDiskStat() (*DiskStat, error) {
	// 参考: disk_test.go
	var path string
	if runtime.GOOS == "windows" {
		path = "C:"
	} else {
		path = "/"
	}
	stat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return &DiskStat{
		UsageStat: stat,
		Path:      path,
	}, nil

	//parts, err := disk.Partitions(true)
	//if err != nil {
	//	return nil, err
	//}
	//for _, part := range parts {
	//	if part.Mountpoint != "/" {
	//		continue
	//	}
	//	stat, err := disk.Usage(part.Mountpoint)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return (*DiskStat)(stat), nil
	//}
	//return nil, errorKit.Newf("fail to get disk stat with parts(length: %d)", len(parts))
}
