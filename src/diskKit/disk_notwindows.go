//go:build !windows

package diskKit

func GetDiskUsageStats() (*DiskUsageStats, error) {
	path := "/"
	return GetDiskUsageStatsByPath(path)
}
