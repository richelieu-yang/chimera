package diskKit

func GetDiskUsageStat() (*DiskUsageStats, error) {
	path := "C:"
	return GetDiskUsageStatsByPath(path)
}
