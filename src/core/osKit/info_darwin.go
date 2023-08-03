package osKit

import (
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"strconv"
)

// GetCurrentCountOfProcesses
/*
支持: 	Linux、Mac
*/
func GetCurrentCountOfProcesses() (int, error) {
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

func GetCurrentCountOfProcessesAndThreads() (int, error) {
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
