package diskKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/shirou/gopsutil/v3/disk"
)

type (
	// DiskUsageStat
	/*
		Total = Used + Free
	*/
	DiskUsageStat struct {
		*disk.UsageStat

		Path string `json:"path"`
	}
)

// GetDiskUsageStatByPath
/*
PS:
(1) Mac（Linux），查看磁盘空间的命令: df -h

golang 获取cpu 内存 硬盘 使用率 信息 进程信息
	https://blog.csdn.net/whatday/article/details/109620192
*/
func GetDiskUsageStatByPath(path string) (*DiskUsageStat, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return &DiskUsageStat{
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
	//	return (*DiskUsageStat)(stat), nil
	//}
	//return nil, errorKit.Newf("fail to get disk stat with parts(length: %d)", len(parts))
}

func (stat *DiskUsageStat) String() string {
	return fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, used percent: %.2f%%",
		stat.Path,
		dataSizeKit.ToReadableStringWithIEC(stat.Free),
		dataSizeKit.ToReadableStringWithIEC(stat.Used),
		dataSizeKit.ToReadableStringWithIEC(stat.Total),
		stat.UsedPercent,
	)
}
