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
func PrintBasicDetails(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	logger.Infof("[CHIMERA] ===================================================================================")

	logger.Infof("[CHIMERA, PROCESS] pid: [%d].", runtimeKit.PID)

	// os
	printOsInfo()

	// golang
	logger.Infof("[CHIMERA, GO] go version: [%s].", runtimeKit.GoVersion)
	logger.Infof("[CHIMERA, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	// user
	logger.Infof("[CHIMERA, USER] name: [%s].", userKit.GetName())
	logger.Infof("[CHIMERA, USER] user name: [%s].", userKit.GetUserName())
	logger.Infof("[CHIMERA, PATH] home dir: [%s].", userKit.GetUserHomeDir())

	// path
	logger.Infof("[CHIMERA, PATH] working directory: [%s].", pathKit.GetWorkingDir())
	logger.Infof("[CHIMERA, PATH] temporary directory: [%s].", pathKit.GetTempDir())
	logger.Infof("[CHIMERA, PATH] SelfDir: [%s].", pathKit.SelfDir())
	logger.Infof("[CHIMERA, PATH] MainPkgPath: [%s].", pathKit.MainPkgPath())

	// time
	systemTime := timeKit.GetSystemTime()
	zoneName, zoneOffset := systemTime.Zone()
	logger.Infof("[CHIMERA, TIME] system time: [%v], zone: [%s, %d].", systemTime, zoneName, zoneOffset)
	if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, TIME] fail to get network time")
	} else {
		logger.Infof("[CHIMERA, TIME] network time: [%v], source: [%s].", networkTime, source)
	}

	// ip
	if ip, err := ipKit.GetOutboundIP(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, IP] fail to get local ip")
	} else {
		logger.Infof("[CHIMERA, IP] local ip(for reference only): [%s].", ip)
	}

	// host
	if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, HOST] fail to get host stat")
	} else {
		logger.Infof("[CHIMERA, HOST] host name: [%s].", hostInfo.Hostname)
	}

	// cpu
	logger.Infof("[CHIMERA, CPU] in virtual machine? [%t].", cpuKit.InVirtualMachine())
	logger.Infof("[CHIMERA, CPU] vendor: [%s].", cpuKit.GetVendor())
	logger.Infof("[CHIMERA, CPU] brand name: [%s].", cpuKit.GetBrandName())
	logger.Infof("[CHIMERA, CPU] number: [%d].", cpuKit.GetNumber())
	if cpuPercent, err := cpuKit.GetPercent(); err != nil {
		logger.Warnf("[CHIMERA, CPU] fail to get cpu percent, error: %v", err)
	} else {
		logger.Infof("[CHIMERA, CPU] usage: [%.2f]%%.", cpuPercent)
	}

	//// mac
	//if macAddresses, err := runtimeKit.GetMacAddresses(); err != nil {
	//	logger.WithFields(logger.Fields{
	//		"error": err.Error(),
	//	}).Fatal("fail to get mac addresses")
	//} else {
	//	logger.Infof("[CHIMERA, MAC] mac addresses: [%v].", macAddresses)
	//}

	// memory
	if stat, err := memoryKit.GetMemoryStat(); err != nil {
		logger.WithError(err).Fatal("[CHIMERA, MEMORY] fail to get memory stat")
	} else {
		logger.Infof("[CHIMERA, MEMORY] stat: [%s].", memoryKit.MemoryStatToString(stat))
	}

	// disk
	if stat, err := diskKit.GetDiskStat(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, DISK] fail to get disk stat")
	} else {
		logger.Infof("[CHIMERA, DISK] stat: [%s].", stat.String())
	}

	// docker
	if dockerIds, err := runtimeKit.GetDockerIdList(); err != nil {
		if err == docker.ErrDockerNotAvailable {
			logger.Info("[CHIMERA, DOCKER] docker isn't available")
		} else {
			logger.WithError(err).Warn("[CHIMERA, DOCKER] Fail to get docker id list")
		}
	} else {
		logger.Infof("[CHIMERA, DOCKER] docker id list: %v.", dockerIds)
	}

	logger.Infof("[CHIMERA] ===================================================================================")
}
