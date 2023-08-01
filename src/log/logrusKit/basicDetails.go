package logrusKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/core/userKit"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
	"github.com/richelieu-yang/chimera/v2/src/ipKit"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/sirupsen/logrus"
)

// PrintBasicDetails 输出服务器的基本信息（以便于甩锅）
func PrintBasicDetails() {
	logrus.Infof("[CHIMERA] ===================================================================================")

	logrus.Infof("[CHIMERA, PROCESS] pid: [%d].", runtimeKit.PID)

	// os
	printOsInfo()

	// golang
	logrus.Infof("[CHIMERA, GO] go version: [%s].", runtimeKit.GoVersion)
	logrus.Infof("[CHIMERA, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	// user
	logrus.Infof("[CHIMERA, USER] name: [%s].", userKit.GetName())
	logrus.Infof("[CHIMERA, USER] user name: [%s].", userKit.GetUserName())
	logrus.Infof("[CHIMERA, PATH] home dir: [%s].", userKit.GetUserHomeDir())

	// path
	logrus.Infof("[CHIMERA, PATH] working directory: [%s].", pathKit.GetWorkingDir())
	logrus.Infof("[CHIMERA, PATH] temporary directory: [%s].", pathKit.GetTempDir())
	logrus.Infof("[CHIMERA, PATH] SelfDir: [%s].", pathKit.SelfDir())
	logrus.Infof("[CHIMERA, PATH] MainPkgPath: [%s].", pathKit.MainPkgPath())

	// time
	systemTime := timeKit.GetSystemTime()
	zoneName, zoneOffset := systemTime.Zone()
	logrus.Infof("[CHIMERA, TIME] system time: [%v], zone: [%s, %d].", systemTime, zoneName, zoneOffset)
	if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, TIME] fail to get network time")
	} else {
		logrus.Infof("[CHIMERA, TIME] network time: [%v], source: [%s].", networkTime, source)
	}

	// ip
	if ip, err := ipKit.GetOutboundIP(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, IP] fail to get local ip")
	} else {
		logrus.Infof("[CHIMERA, IP] local ip(for reference only): [%s].", ip)
	}

	// host
	if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, HOST] fail to get host stat")
	} else {
		logrus.Infof("[CHIMERA, HOST] host name: [%s].", hostInfo.Hostname)
	}

	// cpu
	logrus.Infof("[CHIMERA, CPU] in virtual machine? [%t].", cpuKit.InVirtualMachine())
	logrus.Infof("[CHIMERA, CPU] vendor: [%s].", cpuKit.GetVendor())
	logrus.Infof("[CHIMERA, CPU] brand name: [%s].", cpuKit.GetBrandName())
	logrus.Infof("[CHIMERA, CPU] number: [%d].", cpuKit.GetNumber())
	if cpuPercent, err := cpuKit.GetPercent(); err != nil {
		logrus.Warnf("[CHIMERA, CPU] fail to get cpu percent, error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, CPU] usage: [%.2f]%%.", cpuPercent)
	}

	//// mac
	//if macAddresses, err := runtimeKit.GetMacAddresses(); err != nil {
	//	logrus.WithFields(logrus.Fields{
	//		"error": err.Error(),
	//	}).Fatal("fail to get mac addresses")
	//} else {
	//	logrus.Infof("[CHIMERA, MAC] mac addresses: [%v].", macAddresses)
	//}

	// memory
	if stat, err := memoryKit.GetMemoryStat(); err != nil {
		logrus.WithError(err).Fatal("[CHIMERA, MEMORY] fail to get memory stat")
	} else {
		logrus.Infof("[CHIMERA, MEMORY] stat: [%s].", memoryKit.MemoryStatToString(stat))
	}

	// disk
	if stat, err := diskKit.GetDiskStat(); err != nil {
		logrus.WithError(err).Warn("[CHIMERA, DISK] fail to get disk stat")
	} else {
		logrus.Infof("[CHIMERA, DISK] stat: [%s].", stat.String())
	}

	// docker
	if dockerIds, err := runtimeKit.GetDockerIdList(); err != nil {
		if err == docker.ErrDockerNotAvailable {
			logrus.Info("[CHIMERA, DOCKER] docker isn't available")
		} else {
			logrus.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("[CHIMERA, DOCKER] Fail to get docker id list")
		}
	} else {
		logrus.Infof("[CHIMERA, DOCKER] docker id list: %v.", dockerIds)
	}

	logrus.Infof("[CHIMERA] ===================================================================================")
}
