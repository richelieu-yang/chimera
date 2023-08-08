package diskKit

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type (
	// DiskUsageStats
	/*
		Total = Used + Free
	*/
	DiskUsageStats struct {
		*disk.UsageStat

		Path string `json:"path"`
	}
)

// GetDiskUsageStatsByPath
/*
PS:
(1) Mac（Linux），查看磁盘空间的命令: df -h

golang 获取cpu 内存 硬盘 使用率 信息 进程信息
	https://blog.csdn.net/whatday/article/details/109620192
*/
func GetDiskUsageStatsByPath(path string) (*DiskUsageStats, error) {
	stat, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	return &DiskUsageStats{
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
	//	return (*DiskUsageStats)(stat), nil
	//}
	//return nil, errorKit.Newf("fail to get disk stat with parts(length: %d)", len(parts))
}
