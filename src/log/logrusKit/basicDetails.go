package logrusKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/core/userKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
	"github.com/richelieu-yang/chimera/v2/src/dockerKit"
	"github.com/richelieu-yang/chimera/v2/src/ipKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/sirupsen/logrus"
)

// PrintBasicDetails 输出服务器的基本信息（以便于甩锅）
func PrintBasicDetails(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	logger.Infof("[CHIMERA] ===================================================================================")

	logger.Infof("[CHIMERA, PROCESS] pid: [%d].", processKit.PID)

	// golang
	logger.Infof("[CHIMERA, GO] version: [%s].", runtimeKit.GoVersion)
	logger.Infof("[CHIMERA, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	// os
	printOsInfo()

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
		logger.WithError(err).Warn("[CHIMERA, HOST] fail to get host stats")
	} else {
		logger.Infof("[CHIMERA, HOST] host name: [%s].", hostInfo.Hostname)
	}

	// cpu
	logger.Infof("[CHIMERA, CPU] in a virtual machine? [%t].", cpuKit.InVirtualMachine())
	logger.Infof("[CHIMERA, CPU] family: [%d].", cpuKit.GetFamily())
	logger.Infof("[CHIMERA, CPU] model: [%d].", cpuKit.GetModel())
	logger.Infof("[CHIMERA, CPU] vendor id: [%s].", cpuKit.GetVendorID())
	logger.Infof("[CHIMERA, CPU] vendor string: [%s].", cpuKit.GetVendorString())
	logger.Infof("[CHIMERA, CPU] brand name: [%s].", cpuKit.GetBrandName())
	logger.Infof("[CHIMERA, CPU] physical cores: [%d].", cpuKit.GetPhysicalCores())
	logger.Infof("[CHIMERA, CPU] threads per core: [%d].", cpuKit.GetThreadsPerCore())
	logger.Infof("[CHIMERA, CPU] logical cores: [%d].", cpuKit.GetLogicalCores())
	logger.Infof("[CHIMERA, CPU] CPU number: [%d].", cpuKit.GetCpuNumber())
	logger.Infof("[CHIMERA, CPU] features: [%s].", sliceKit.Join(cpuKit.GetFeatureSet(), ","))
	logger.Infof("[CHIMERA, CPU] frequency: [%d]hz.", cpuKit.GetFrequency())
	if cpuPercent, err := cpuKit.GetUsage(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, CPU] fail to get cpu usage")
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
	if stats, err := memoryKit.GetMachineMemoryStats(); err != nil {
		logger.WithError(err).Fatal("[CHIMERA, MEMORY] fail to get machine memory stats")
	} else {
		str := fmt.Sprintf("total: %s, available: %s, used: %s, free: %s, used percent: %.2f%%",
			dataSizeKit.ToReadableStringWithIEC(stats.Total),
			dataSizeKit.ToReadableStringWithIEC(stats.Available),
			dataSizeKit.ToReadableStringWithIEC(stats.Used),
			dataSizeKit.ToReadableStringWithIEC(stats.Free),
			stats.UsedPercent,
		)
		logger.Infof("[CHIMERA, MEMORY] machine memory stats: [%s].", str)
	}

	// disk
	if stats, err := diskKit.GetDiskUsageStat(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, DISK] fail to get disk usage stats")
	} else {
		str := fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, used percent: %.2f%%",
			stats.Path,
			dataSizeKit.ToReadableStringWithIEC(stats.Free),
			dataSizeKit.ToReadableStringWithIEC(stats.Used),
			dataSizeKit.ToReadableStringWithIEC(stats.Total),
			stats.UsedPercent,
		)
		logger.Infof("[CHIMERA, DISK] disk usage stats: [%s].", str)
	}

	// docker
	if dockerIds, err := dockerKit.GetDockerIdList(); err != nil {
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
