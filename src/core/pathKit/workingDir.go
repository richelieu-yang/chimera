package pathKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

// GetCurrentWorkingDir 获取 当前工作目录的绝对路径
func GetCurrentWorkingDir() (string, error) {
	return os.Getwd()
}

// ChangeWorkingDir 设置 当前工作目录的绝对路径
func ChangeWorkingDir(dir string) error {
	return gfile.Chdir(dir)
}
