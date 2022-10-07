package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

func GetCpuNumber() int {
	return runtime.NumCPU()
}

// GetCpuPercent CPU使用率
func GetCpuPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}
