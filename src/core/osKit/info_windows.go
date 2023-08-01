package osKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

func GetCountOfProcesses() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetMaxOpenFiles() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetMaxUserProcesses() (int, error) {
	return 0, errorKit.New("not yet realized")
}

func GetCoreFileSize() (string, error) {
	return "", errorKit.New("not yet realized")
}
