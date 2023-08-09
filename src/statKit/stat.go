package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/floatKit"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

type (
	Stats struct {
		Cpu *CpuStats `json:"cpu"`

		Disk *DiskStats `json:"disk"`

		Program *ProgramStats `json:"program"`

		Machine *MachineStats `json:"machine"`
	}

	CpuStats struct {
		Usage      float64 `json:"usage,omitempty"`
		UsageError error   `json:"usageError,omitempty"`
	}

	DiskStats struct {
		Path       string  `json:"path,omitempty"`
		Usage      float64 `json:"usage,omitempty"`
		UsageError error   `json:"usageError,omitempty"`
	}

	ProgramStats struct {
		GoroutineCount int `json:"goroutineCount"`

		Alloc      string `json:"alloc"`
		TotalAlloc string `json:"totalAlloc"`
		Sys        string `json:"sys"`
		NumGC      uint32 `json:"numGC"`
		EnableGC   bool   `json:"enableGC"`
	}

	MachineStats struct {
		// ProcessCount 进程数
		ProcessCount      int   `json:"processCount,omitempty"`
		ProcessCountError error `json:"processCountError,omitempty"`

		// ProcessThreadCount 进程数（包括线程数）
		ProcessThreadCount      int   `json:"processThreadCount,omitempty"`
		ProcessThreadCountError error `json:"processThreadCountError,omitempty"`

		MemoryStatsError error   `json:"memoryStatsError,omitempty"`
		Total            string  `json:"total,omitempty"`
		Available        string  `json:"available,omitempty"`
		Used             string  `json:"used,omitempty"`
		UsedPercent      float64 `json:"usedPercent,omitempty"`
		Free             string  `json:"free,omitempty"`
	}
)

// GetStats
/*
PS: 由于获取CPU使用率耗时较长，使用 sync.WaitGroup.
*/
func GetStats() (rst *Stats) {
	rst = &Stats{}
	var wg sync.WaitGroup

	/* CPU */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var cpuStats = &CpuStats{}
		rst.Cpu = cpuStats
		{
			usage, err := cpuKit.GetUsage()
			if err != nil {
				cpuStats.UsageError = err
			} else {
				cpuStats.Usage = floatKit.Round(usage, 2)
			}
		}
	}()

	/* DISK */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var diskStats = &DiskStats{}
		rst.Disk = diskStats
		{
			stats, err := diskKit.GetDiskUsageStat()
			if err != nil {
				diskStats.UsageError = err
			} else {
				diskStats.Path = stats.Path
				diskStats.Usage = floatKit.Round(stats.UsedPercent, 2)
			}
		}
	}()

	/* program */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var programStats = &ProgramStats{}
		rst.Program = programStats
		{
			stats := memoryKit.GetProgramMemoryStats()

			programStats.GoroutineCount = runtime.NumGoroutine()

			programStats.Alloc = dataSizeKit.ToReadableStringWithIEC(stats.Alloc)
			programStats.TotalAlloc = dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc)
			programStats.Sys = dataSizeKit.ToReadableStringWithIEC(stats.Sys)
			programStats.NumGC = stats.NumGC
			programStats.EnableGC = stats.EnableGC
		}
	}()

	/* machine */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var machineStats = &MachineStats{}
		rst.Machine = machineStats
		{
			count, err := processKit.GetProcessCount()
			if err != nil {
				machineStats.ProcessCountError = err
			} else {
				machineStats.ProcessCount = count
			}

			count1, err := processKit.GetProcessThreadCount()
			if err != nil {
				machineStats.ProcessThreadCountError = err
			} else {
				machineStats.ProcessThreadCount = count1
			}

			stats, err := memoryKit.GetMachineMemoryStats()
			if err != nil {
				machineStats.MemoryStatsError = err
			} else {
				machineStats.Total = dataSizeKit.ToReadableStringWithIEC(stats.Total)
				machineStats.Available = dataSizeKit.ToReadableStringWithIEC(stats.Available)
				machineStats.Used = dataSizeKit.ToReadableStringWithIEC(stats.Used)
				machineStats.UsedPercent = floatKit.Round(stats.UsedPercent, 2)
				machineStats.Free = dataSizeKit.ToReadableStringWithIEC(stats.Free)
			}
		}
	}()

	wg.Wait()
	return rst
}

func PrintStats(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	stats := GetStats()
	json, err := jsonKit.MarshalIndentToString(stats, "", "    ")
	if err != nil {
		logger.WithError(err).Error("fail to print")
	}
	logrusKit.DisableQuoteTemporarily(logger, func(logger *logrus.Logger) {
		logger.Infof("[CHIMERA] stats:\n%s", json)
	})
}
