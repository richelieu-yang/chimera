package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"strconv"
)

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

// GetThreadsMax 获取"kernel.threads-max"
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

// GetPidMax 获取"kernel.pid_max"
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
