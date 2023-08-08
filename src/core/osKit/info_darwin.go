package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"strconv"
)

// GetProcessCount 获取: (瞬时的值)系统中所有进程的数量.
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

// GetProcessThreadCount 获取: (瞬时的值)系统中所有进程及其线程的数量.
func GetProcessThreadCount() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetThreadsMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetPidMax() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetMaxMapCount() (int, error) {
	return 0, errorKit.New("not yet realized")
}
