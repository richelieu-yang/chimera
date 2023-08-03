//go:build linux || darwin

package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/sirupsen/logrus"
)

// printUniqueOsInfo 输出特殊的信息（主要依赖于不同的OS）
func printUniqueOsInfo() {
	if count, err := osKit.GetCountOfProcesses(); err != nil {
		logrus.WithError(err).Error("[CHIMERA, OS] fail to get current count of processes")
	} else {
		logrus.Infof("[CHIMERA, OS] current count of processes: [%d].", count)
	}

	if str, err := osKit.GetUlimitInfo(); err != nil {
		logrus.WithError(err).Error("[CHIMERA, OS] fail to get ulimit information")
	} else {
		DisableQuoteTemporarily(nil, func() {
			logrus.Infof("[CHIMERA, OS] ulimit information:\n%s", str)
		})
	}

	//if userMaxProcesses, err := osKit.GetMaxUserProcesses(); err != nil {
	//	logrus.WithError(err).Error("[CHIMERA, OS] fail to get ulimit -u(max user processes)")
	//} else {
	//	logrus.Infof("[CHIMERA, OS] ulimit -u(max user processes): [%d].", userMaxProcesses)
	//}
	//
	//if maxOpenFiles, err := osKit.GetMaxOpenFiles(); err != nil {
	//	logrus.WithError(err).Error("[CHIMERA, OS] fail to get ulimit -n(open files)")
	//} else {
	//	logrus.Infof("[CHIMERA, OS] ulimit -n(open files): [%d].", maxOpenFiles)
	//}
	//
	//if coreFileSize, err := osKit.GetCoreFileSize(); err != nil {
	//	logrus.WithError(err).Error("[CHIMERA, OS] fail to get ulimit -c(core file size)")
	//} else {
	//	logrus.Infof("[CHIMERA, OS] ulimit -c(core file size): [%s].", coreFileSize)
	//}
}
