package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
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

func GetThreadCount() (int, error) {
	return 0, errorKit.New("not yet realized")
}

// GetThreadsMax 获取Linux的"kernel.threads-max"
func GetThreadsMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}

// GetPidMax 获取Linux的"kernel.pid_max"
func GetPidMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}

// GetMaxThreadCountInAProcess 获取 Linux的"vm.max_map_count"（单进程可生成的最大线程数）
func GetMaxThreadCountInAProcess() (int, error) {
	return 0, errorKit.New("not yet realized")
}
