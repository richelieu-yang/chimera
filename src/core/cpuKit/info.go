package cpuKit

import (
	"github.com/klauspost/cpuid/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"runtime"
	"time"
)

// GetNumber returns the number of logical CPUs usable by the current process.
func GetNumber() int {
	return runtime.NumCPU()
}

// GetPercent CPU使用率
func GetPercent() (float64, error) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	return percent[0], nil
}

// GetVendor CPU供应商
/*
@return e.g."Apple"
*/
func GetVendor() string {
	return cpuid.CPU.VendorString
}

// GetBrandName CPU品牌名称
/*
@return e.g."Apple M1 Pro"
*/
func GetBrandName() string {
	return cpuid.CPU.BrandName
}
