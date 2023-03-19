package pathKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	var err error

	// projectDir
	projectDir, err = os.Getwd()
	if err != nil {
		logrus.Panicf("[SCALES, PATH] Funtion(os.Getwd()) fails, error: %v", err)
	}
	if strKit.IsEmpty(projectDir) {
		logrus.Panic("[SCALES, PATH] Variable projectDir is empty.")
	}
	if !fileKit.IsDir(projectDir) {
		logrus.Panicf("[SCALES, PATH] Variable projectDir(%s) isn't a directory.", projectDir)
	}
}
