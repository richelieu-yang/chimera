package logrusKit

import (
	"errors"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/userKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
	"github.com/richelieu-yang/chimera/v2/src/dockerKit"
	"github.com/richelieu-yang/chimera/v2/src/ip/ipKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"github.com/richelieu-yang/chimera/v2/src/serialize/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/time/timeKit"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/sirupsen/logrus"
)

// PrintBasicDetails 输出服务器的基本信息（以便于甩锅）
func PrintBasicDetails(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}
	out := logger.Out

	_, _ = out.Write([]byte("===========================================================\n"))
	_, _ = out.Write([]byte(consts.Banner))

	logger.Infof("[CHIMERA, PROCESS] pid: [%d].", processKit.PID)

	/* golang */
	logger.Infof("[CHIMERA, GO] version: [%s].", runtimeKit.GoVersion)
	logger.Infof("[CHIMERA, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	/* os */
	printOsInfo()

	/* user */
	logger.Infof("[CHIMERA, USER] name: [%s].", userKit.GetName())
	logger.Infof("[CHIMERA, USER] user name: [%s].", userKit.GetUserName())
	logger.Infof("[CHIMERA, PATH] home dir: [%s].", userKit.GetUserHomeDir())

	/* path */
	logger.Infof("[CHIMERA, PATH] working directory: [%s].", pathKit.GetWorkingDir())
	logger.Infof("[CHIMERA, PATH] temporary directory: [%s].", pathKit.GetTempDir())
	logger.Infof("[CHIMERA, PATH] SelfDir: [%s].", pathKit.SelfDir())
	logger.Infof("[CHIMERA, PATH] MainPkgPath: [%s].", pathKit.MainPkgPath())

	/* json */
	logger.Infof("[CHIMERA, JSON] library: [%s].", jsonKit.GetLibrary())

	/* time */
	machineTime := timeKit.GetMachineTime()
	zoneName, zoneOffset := machineTime.Zone()
	logger.Infof("[CHIMERA, TIME] machine time: [%v], zone: [%s, %d].", machineTime, zoneName, zoneOffset)
	// Richelieu: 先注释掉，以防（断网或弱网环境下）导致拖延服务启动3s
	//if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
	//	logger.WithError(err).Warn("[CHIMERA, TIME] fail to get network time")
	//} else {
	//	logger.Infof("[CHIMERA, TIME] network time: [%v], source: [%s].", networkTime, source)
	//}

	/* ip */
	InternalIp := ipKit.GetInternalIp()
	logger.Infof("[CHIMERA, IP] internal ip: [%s].", InternalIp)
	ips := ipKit.GetIps()
	logger.Infof("[CHIMERA, IP] ips: [%s].", sliceKit.Join(ips, ", "))

	/* host */
	if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, HOST] Fail to get host info.")
	} else {
		logger.Infof("[CHIMERA, HOST] host name: [%s].", hostInfo.Hostname)
	}

	/* cpu */
	cpuKit.PrintBasicDetails(logger)

	/* mac */
	//if macAddresses, err := runtimeKit.GetMacAddresses(); err != nil {
	//	logger.WithFields(logger.Fields{
	//		"error": err.Error(),
	//	}).Fatal("fail to get mac addresses")
	//} else {
	//	logger.Infof("[CHIMERA, MAC] mac addresses: [%v].", macAddresses)
	//}

	/* memory */
	stats, err := memoryKit.GetMachineMemoryStats()
	if err != nil {
		logger.WithError(err).Fatal("[CHIMERA, MEMORY] Fail to get machine memory stats.")
		return
	}
	str := fmt.Sprintf("total: %s, available: %s, used: %s, free: %s, used percent: %.2f%%",
		dataSizeKit.ToReadableIecString(float64(stats.Total)),
		dataSizeKit.ToReadableIecString(float64(stats.Available)),
		dataSizeKit.ToReadableIecString(float64(stats.Used)),
		dataSizeKit.ToReadableIecString(float64(stats.Free)),
		stats.UsedPercent,
	)
	logger.Infof("[CHIMERA, MEMORY] machine memory stats: [%s].", str)

	/* disk */
	diskKit.PrintBasicDetails(logger)

	//if stats, err := diskKit.GetDiskUsageStats(); err != nil {
	//	logger.WithError(err).Warn("[CHIMERA, DISK] Fail to get disk usage stats.")
	//} else {
	//	str := fmt.Sprintf("path: %s, free: %s, used: %s, total: %s, used percent: %.2f%%",
	//		stats.Path,
	//		dataSizeKit.ToReadableIecString(float64(stats.Free)),
	//		dataSizeKit.ToReadableIecString(float64(stats.Used)),
	//		dataSizeKit.ToReadableIecString(float64(stats.Total)),
	//		stats.UsedPercent,
	//	)
	//	logger.Infof("[CHIMERA, DISK] disk usage stats: [%s].", str)
	//}

	// Richelieu: 要有条件地使用 gopsutil
	/* docker */
	if dockerIds, err := dockerKit.GetDockerIdList(); err != nil {
		if errors.Is(err, docker.ErrDockerNotAvailable) {
			logger.Info("[CHIMERA, DOCKER] Docker isn't available.")
		} else {
			logger.WithError(err).Warn("[CHIMERA, DOCKER] Fail to get docker id list.")
		}
	} else {
		logger.Infof("[CHIMERA, DOCKER] docker id list: %v.", dockerIds)
	}

	_, _ = out.Write([]byte("===========================================================\n"))
}
