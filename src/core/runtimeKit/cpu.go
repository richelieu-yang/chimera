package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

// GetCPUNumber returns the number of logical CPUs usable by the current process.
func GetCPUNumber() int {
	return runtime.NumCPU()
}

// GetCPUPercent CPU使用率
func GetCPUPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}
