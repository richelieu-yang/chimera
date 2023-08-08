//go:build !windows

package diskKit

func GetDiskUsageStat() (*DiskUsageStat, error) {
	path := "/"
	return GetDiskUsageStatByPath(path)
}

//var GetUsageStat func(path string) (*disk.UsageStat, error) = disk.Usage
//
//// GetDiskUsageStat
///*
//PS:
//(1) Mac（Linux），查看磁盘空间的命令: df -h
//
//golang 获取cpu 内存 硬盘 使用率 信息 进程信息
//	https://blog.csdn.net/whatday/article/details/109620192
//*/
//func GetDiskUsageStat() (*DiskUsageStat, error) {
//	// 参考: disk_test.go
//	var path string
//	if runtime.GOOS == "windows" {
//		path = "C:"
//	} else {
//		path = "/"
//	}
//
//	stat, err := GetUsageStat(path)
//	if err != nil {
//		return nil, err
//	}
//	return &DiskUsageStat{
//		UsageStat: stat,
//		Path:      path,
//	}, nil
//
//	//parts, err := disk.Partitions(true)
//	//if err != nil {
//	//	return nil, err
//	//}
//	//for _, part := range parts {
//	//	if part.Mountpoint != "/" {
//	//		continue
//	//	}
//	//	stat, err := disk.Usage(part.Mountpoint)
//	//	if err != nil {
//	//		return nil, err
//	//	}
//	//	return (*DiskUsageStat)(stat), nil
//	//}
//	//return nil, errorKit.Newf("fail to get disk stat with parts(length: %d)", len(parts))
//}
