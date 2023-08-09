//go:build !windows

package diskKit

func GetDiskUsageStat() (*DiskUsageStats, error) {
	path := "/"
	return GetDiskUsageStatsByPath(path)
}
