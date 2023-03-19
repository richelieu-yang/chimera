package pathKit

import (
	"github.com/richelieu42/chimera/src/consts"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/sirupsen/logrus"
)

// projectDir 项目目录（绝对路径）
var projectDir string

func GetProjectDir() string {
	return projectDir
}

func SetProjectDir(dir string) error {
	if err := fileKit.MkDirs(dir); err != nil {
		return err
	}
	if err := fileKit.AssertExistAndIsDir(dir); err != nil {
		return err
	}
	projectDir = dir
	return nil
}

// ReviseProjectDirWhenTesting 在测试(执行xxx_test.go文件)时，自动修改项目路径，将 projectDir 修改为 当前项目的根目录.
/*
PS:
(1) 执行xxx_test.go文件时，projectDir为该文件所在的目录，而非当前项目的根目录.
(2) 想调用此方法的话，必须在程序启动时立即调用，以防来不及.
*/
func ReviseProjectDirWhenTesting() error {
	old := projectDir

	index := strKit.LastIndex(old, consts.Name)
	if index == -1 {
		return errorKit.Simple("strKit.LastIndex(projectDir, consts.Name) == -1")
	}
	tmp := strKit.SubBefore(old, index+len(consts.Name))
	if err := SetProjectDir(tmp); err != nil {
		return err
	}
	logrus.Infof("[SCALES, PATH] In test mode, revise projectDir from [%s] to [%s].", old, projectDir)
	return nil
}
