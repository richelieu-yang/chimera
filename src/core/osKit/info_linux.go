package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"strconv"
)

// GetProcessCount
/*
支持: 	Linux、Mac
*/
func GetProcessCount() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ps auxw | wc -l")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetThreadCount
/*
支持: 	Linux
不支持:	Mac
*/
func GetThreadCount() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ps -eLf | wc -l")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetThreadsMax 获取 Linux的"kernel.threads-max"（系统可生成最大线程数）
/*
命令:
cat /proc/sys/kernel/threads-max
sysctl kernel.threads-max
*/
func GetThreadsMax() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "cat /proc/sys/kernel/threads-max")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetPidMax 获取 Linux的"kernel.pid_max"
/*
命令:
cat /proc/sys/kernel/pid_max
sysctl kernel.pid_max
*/
func GetPidMax() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "cat /proc/sys/kernel/pid_max")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// GetMaxThreadCountInAProcess 获取 Linux的"vm.max_map_count"（单进程可生成的最大线程数）
/*
命令:
cat /proc/sys/vm/max_map_count
sysctl vm.max_map_count
*/
func GetMaxThreadCountInAProcess() (int, error) {
	str, err := cmdKit.ExecuteToString("sh", "-c", "cat /proc/sys/vm/max_map_count")
	if err != nil {
		return 0, err
	}
	str = strKit.TrimSpace(str)

	i, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return i, nil
}
