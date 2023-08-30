package diskKit

func GetDiskUsageStats() (*DiskUsageStats, error) {
	path := "C:"
	return GetDiskUsageStatsByPath(path)
}
