package pathKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/sirupsen/logrus"
	"os"
)

// GetWorkingDir 获取 当前工作目录的绝对路径
func GetWorkingDir() string {
	path, err := os.Getwd()
	if err != nil {
		logrus.WithError(err).Fatal("fail to get working directory")
	}
	return path
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
