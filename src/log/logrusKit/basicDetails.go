package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/osKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/runtimeKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/richelieu42/chimera/src/core/userKit"
	"github.com/richelieu42/chimera/src/ipKit"
	"github.com/sirupsen/logrus"
)

// PrintBasicDetails 输出服务器的基本信息（以便于甩锅）
func PrintBasicDetails() {
	logrus.Infof("[CHIMERA] ======================================================================================================================")

	logrus.Infof("[CHIMERA, PROCESS] pid: [%d].", runtimeKit.PID)

	// os
	logrus.Infof("[CHIMERA, OS] os: [%s].", osKit.OS)
	logrus.Infof("[CHIMERA, OS] arch: [%s].", osKit.ARCH)
	logrus.Infof("[CHIMERA, OS] bits: [%d].", osKit.BITS)
	printUniqueOsInfo()

	// golang
	logrus.Infof("[CHIMERA, GO] go version: [%s].", runtimeKit.GoVersion)
	logrus.Infof("[CHIMERA, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	// user
	logrus.Infof("[CHIMERA, USER] name: [%s].", userKit.GetName())
	logrus.Infof("[CHIMERA, USER] user name: [%s].", userKit.GetUserName())
	logrus.Infof("[CHIMERA, PATH] home dir: [%s].", userKit.GetUserHomeDir())

	// path
	logrus.Infof("[CHIMERA, PATH] project path: [%s].", pathKit.GetProjectDir())
	logrus.Infof("[CHIMERA, PATH] temporary directory: [%s].", pathKit.GetTempDir())

	// time
	systemTime := timeKit.GetSystemTime()
	zoneName, zoneOffset := systemTime.Zone()
	logrus.Infof("[CHIMERA, TIME] system time: [%v], zone: [%s, %d].", systemTime, zoneName, zoneOffset)
	if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
		logrus.Warnf("[CHIMERA, TIME] fail to get network time, error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, TIME] network time: [%v], source: [%s].", networkTime, source)
	}

	// ip
	if ip, err := ipKit.GetOutboundIP(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("[CHIMERA, IP] fail to get local ip")
	} else {
		logrus.Infof("[CHIMERA, IP] local ip(for reference only): [%s].", ip)
	}

	// host
	if hostInfo, err := runtimeKit.GetHostInfo(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("[CHIMERA, HOST] fail to get host info")
	} else {
		logrus.Infof("[CHIMERA, HOST] host name: [%s].", hostInfo.Hostname)
	}

	// cpu
	if cpuId, err := runtimeKit.GetCpuId(); err != nil {
		logrus.Warnf("[CHIMERA, CPU] fail to get cpu id, error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, CPU] cpu id: [%s].", cpuId)
	}
	logrus.Infof("[CHIMERA, CPU] available processors: [%d].", runtimeKit.GetCpuNumber())
	if cpuPercent, err := runtimeKit.GetCpuPercent(); err != nil {
		logrus.Warnf("[CHIMERA, CPU] fail to get cpu percent, error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, CPU] usage: [%.2f]%%.", cpuPercent)
	}

	// mac
	if macAddresses, err := runtimeKit.GetMacAddresses(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get mac addresses")
	} else {
		logrus.Infof("[CHIMERA, MAC] mac addresses: [%v].", macAddresses)
	}

	// memory
	if info, err := runtimeKit.GetMemoryStat(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get memory stat")
	} else {
		logrus.Infof("[CHIMERA, MEMORY] memory stat: [%s].", info)
	}

	// docker
	if dockerIds, err := runtimeKit.GetDockerIdList(); err != nil {
		logrus.Warnf("[CHIMERA, DOCKER] Fail to get docker id list, error: %v", err)
	} else {
		logrus.Infof("[CHIMERA, DOCKER] docker id list: %v.", dockerIds)
	}

	logrus.Infof("[CHIMERA] ======================================================================================================================")
}
