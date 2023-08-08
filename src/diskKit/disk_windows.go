package diskKit

func GetDiskUsageStat() (*DiskUsageStat, error) {
	path := "C:"
	return GetDiskUsageStatByPath(path)
}
