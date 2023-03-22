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
	logrus.Infof("[SCALES] ======================================================================================================================")

	logrus.Infof("[SCALES, PROCESS] pid: [%d].", runtimeKit.PID)

	// os
	logrus.Infof("[SCALES, OS] os: [%s].", osKit.OS)
	logrus.Infof("[SCALES, OS] arch: [%s].", osKit.ARCH)
	logrus.Infof("[SCALES, OS] bits: [%d].", osKit.BITS)
	printUniqueOsInfo()

	// golang
	logrus.Infof("[SCALES, GO] go version: [%s].", runtimeKit.GoVersion)
	logrus.Infof("[SCALES, GO] GOROOT: [%s].", runtimeKit.GoRoot)

	// user
	logrus.Infof("[SCALES, USER] name: [%s].", userKit.GetName())
	logrus.Infof("[SCALES, USER] user name: [%s].", userKit.GetUserName())
	logrus.Infof("[SCALES, PATH] home dir: [%s].", userKit.GetUserHomeDir())

	// path
	logrus.Infof("[SCALES, PATH] project path: [%s].", pathKit.GetProjectDir())
	logrus.Infof("[SCALES, PATH] temporary directory: [%s].", pathKit.GetTempDir())

	// time
	systemTime := timeKit.GetSystemTime()
	zoneName, zoneOffset := systemTime.Zone()
	logrus.Infof("[SCALES, TIME] system time: [%v], zone: [%s, %d].", systemTime, zoneName, zoneOffset)
	if networkTime, source, err := timeKit.GetNetworkTime(); err != nil {
		logrus.Warnf("[SCALES, TIME] fail to get network time, error: %v", err)
	} else {
		logrus.Infof("[SCALES, TIME] network time: [%v], source: [%s].", networkTime, source)
	}

	// ip
	if ip, err := ipKit.GetOutboundIP(); err != nil {
		logrus.Warnf("[SCALES, IP] fail to get local ip, error: %v", err)
	} else {
		logrus.Infof("[SCALES, IP] local ip(for reference only): [%s].", ip)

		// ip2region
		if region, err := ipKit.GetRegionByIp(ip); err != nil {
			logrus.Warnf("[SCALES, IP] fail to get region of local ip, error: %v", err)
		} else {
			logrus.Infof("[SCALES, IP] region: [%s].", region)
		}
	}

	// host
	hostInfo := runtimeKit.GetHostInfo()
	logrus.Infof("[SCALES, HOST] host name: [%s].", hostInfo.Hostname)

	// cpu
	if cpuId, err := runtimeKit.GetCpuId(); err != nil {
		logrus.Warnf("[SCALES, CPU] fail to get cpu id, error: %v", err)
	} else {
		logrus.Infof("[SCALES, CPU] cpu id: [%s].", cpuId)
	}
	logrus.Infof("[SCALES, CPU] available processors: [%d].", runtimeKit.GetCpuNumber())
	if cpuPercent, err := runtimeKit.GetCpuPercent(); err != nil {
		logrus.Warnf("[SCALES, CPU] fail to get cpu percent, error: %v", err)
	} else {
		logrus.Infof("[SCALES, CPU] usage: [%.2f]%%.", cpuPercent)
	}

	// mac
	if macAddresses, err := runtimeKit.GetMacAddresses(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get mac addresses")
	} else {
		logrus.Infof("[SCALES, MAC] mac addresses: [%v].", macAddresses)
	}

	//// disk
	//if stat, err := runtimeKit.GetDiskStat(); err != nil {
	//	errorKit.Panic("fail to get disk stat, error:\n%+v", err)
	//} else {
	//	logrus.Infof("[SCALES, DISK] disk stat: [%s].", stat.String())
	//}

	// memory
	if info, err := runtimeKit.GetMemoryStat(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get memory stat")
	} else {
		logrus.Infof("[SCALES, MEMORY] memory stat: [%s].", info)
	}

	// docker
	if dockerIds, err := runtimeKit.GetDockerIdList(); err != nil {
		logrus.Warnf("[SCALES, DOCKER] Fail to get docker id list, error: %v", err)
	} else {
		logrus.Infof("[SCALES, DOCKER] docker id list: %v.", dockerIds)
	}

	logrus.Infof("[SCALES] ======================================================================================================================")
}
