package pathKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"os"
)

// GetWorkingDir 获取 当前工作目录的绝对路径
func GetWorkingDir() (string, error) {
	return os.Getwd()
}

// ChangeWorkingDir 设置 当前工作目录的绝对路径
func ChangeWorkingDir(dir string) error {
	if err := fileKit.AssertNotExistOrIsDir(dir); err != nil {
		return err
	}
	if err := fileKit.MkDirs(dir); err != nil {
		return err
	}

	return gfile.Chdir(dir)
}
