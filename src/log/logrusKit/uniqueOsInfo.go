//go:build !windows

package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/osKit"
	"github.com/sirupsen/logrus"
)

// printUniqueOsInfo 输出特殊的信息（主要依赖于不同的OS）
func printUniqueOsInfo() {
	if userMaxProcesses, err := osKit.GetUserMaxProcesses(); err != nil {
		logrus.Warnf("[CHIMERA, OS] fail to get userMaxProcesses(ulimit -u), error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, OS] userMaxProcesses(ulimit -u): [%d].", userMaxProcesses)
	}

	if maxOpenFiles, err := osKit.GetMaxOpenFiles(); err != nil {
		logrus.Warnf("[CHIMERA, OS] fail to get maxOpenFiles(ulimit -n), error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, OS] maxOpenFiles(ulimit -n): [%d].", maxOpenFiles)
	}
}
