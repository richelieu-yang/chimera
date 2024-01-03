//go:build darwin || windows || (linux && 386) || (linux && amd64) || (linux && arm) || (linux && arm64)

package cpuKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func PrintBasicDetails(logger *logrus.Logger) {
	logger.Infof("[CHIMERA, CPU] in a virtual machine? [%t].", InVirtualMachine())
	logger.Infof("[CHIMERA, CPU] vendor id: [%s].", GetVendorID())
	logger.Infof("[CHIMERA, CPU] vendor string: [%s].", GetVendorString())
	logger.Infof("[CHIMERA, CPU] brand name: [%s].", GetBrandName())
	logger.Infof("[CHIMERA, CPU] CPU number: [%d].", GetCpuNumber())
	logger.Infof("[CHIMERA, CPU] features: [%s].", sliceKit.Join(GetFeatureSet(), ","))
	logger.Infof("[CHIMERA, CPU] frequency: [%d]hz.", GetFrequency())
	if cpuPercent, err := GetUsagePercent(); err != nil {
		logger.WithError(err).Warn("[CHIMERA, CPU] fail to get cpu usage")
	} else {
		logger.Infof("[CHIMERA, CPU] usage percent: [%.2f]%%.", cpuPercent)
	}
}

// GetUsagePercent CPU使用率
/*
PS: 耗时约1s.

e.g.
() => 12.701612903175233
*/
func GetUsagePercent() (float64, error) {
	s, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return s[0], nil
}

// GetCurrentProcessUsagePercent 获取 当前进程 的CPU使用百分比.
/*
PS: 类似Linux命令: top -p ${pid}
*/
func GetCurrentProcessUsagePercent() (float64, error) {
	var pid = int32(os.Getpid())
	return GetProcessUsagePercent(pid)
}

// GetProcessUsagePercent 获取 指定进程 的CPU使用百分比.
func GetProcessUsagePercent(pid int32) (float64, error) {
	p, err := process.NewProcess(pid)
	if err != nil {
		return 0, err
	}
	return p.CPUPercent()
}
