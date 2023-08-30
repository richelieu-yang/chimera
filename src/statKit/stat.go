package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"github.com/sirupsen/logrus"
	"os"
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
		PID            int `json:"pid"`
		GoroutineCount int `json:"goroutineCount"`

		Alloc      string `json:"alloc"`
		TotalAlloc string `json:"totalAlloc"`
		Sys        string `json:"sys"`
		NumGC      uint32 `json:"numGC"`
		EnableGC   bool   `json:"enableGC"`

		CpuUsage      float64 `json:"cpuUsage"`
		CpuUsageError error   `json:"cpuUsageError,omitempty"`
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

		MaxProcessThreadCountByUser      int    `json:"maxProcessThreadCountByUser,omitempty"`
		MaxProcessThreadCountByUserError string `json:"maxProcessThreadCountByUserError,omitempty"`
		PidMax                           int    `json:"pidMax,omitempty"`
		PidMaxError                      string `json:"pidMaxError,omitempty"`
		ThreadsMax                       int    `json:"threadsMax,omitempty"`
		ThreadsMaxError                  string `json:"threadsMaxError,omitempty"`
		MaxMapCount                      int    `json:"maxMapCount,omitempty"`
		MaxMapCountError                 string `json:"maxMapCountError,omitempty"`
	}
)

// GetStats
/*
PS: 由于获取CPU使用率耗时较长，本函数内部使用 sync.WaitGroup.
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
			usage, err := cpuKit.GetUsagePercent()
			if err != nil {
				cpuStats.UsageError = err
			} else {
				cpuStats.Usage = mathKit.Round(usage, 2)
			}
		}
	}()

	/* disk */
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
				diskStats.Usage = mathKit.Round(stats.UsedPercent, 2)
			}
		}
	}()

	/* program */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var pStats = &ProgramStats{}
		rst.Program = pStats
		{
			pStats.PID = os.Getpid()
			pStats.GoroutineCount = runtime.NumGoroutine()

			stats := memoryKit.GetProgramMemoryStats()
			pStats.Alloc = dataSizeKit.ToReadableStringWithIEC(stats.Alloc)
			pStats.TotalAlloc = dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc)
			pStats.Sys = dataSizeKit.ToReadableStringWithIEC(stats.Sys)
			pStats.NumGC = stats.NumGC
			pStats.EnableGC = stats.EnableGC

			if usage, err := cpuKit.GetUsagePercentByProcess(int32(pStats.PID)); err != nil {
				pStats.CpuUsageError = err
			} else {
				pStats.CpuUsage = usage
			}
		}
	}()

	/* machine */
	wg.Add(1)
	go func() {
		defer wg.Done()

		var mStats = &MachineStats{}
		rst.Machine = mStats
		{
			count, err := processKit.GetProcessCount()
			if err != nil {
				mStats.ProcessCountError = err
			} else {
				mStats.ProcessCount = count
			}

			count1, err := processKit.GetProcessThreadCount()
			if err != nil {
				mStats.ProcessThreadCountError = err
			} else {
				mStats.ProcessThreadCount = count1
			}

			stats, err := memoryKit.GetMachineMemoryStats()
			if err != nil {
				mStats.MemoryStatsError = err
			} else {
				mStats.Total = dataSizeKit.ToReadableStringWithIEC(stats.Total)
				mStats.Available = dataSizeKit.ToReadableStringWithIEC(stats.Available)
				mStats.Used = dataSizeKit.ToReadableStringWithIEC(stats.Used)
				mStats.UsedPercent = mathKit.Round(stats.UsedPercent, 2)
				mStats.Free = dataSizeKit.ToReadableStringWithIEC(stats.Free)
			}

			// ulimit -u
			if tmp, err := osKit.GetMaxProcessThreadCountByUser(); err != nil {
				mStats.MaxProcessThreadCountByUserError = err.Error()
			} else {
				mStats.MaxProcessThreadCountByUser = tmp
			}
			// kernel.pid_max
			if tmp, err := osKit.GetPidMax(); err != nil {
				mStats.PidMaxError = err.Error()
			} else {
				mStats.PidMax = tmp
			}
			// kernel.threads-max
			if tmp, err := osKit.GetThreadsMax(); err != nil {
				mStats.ThreadsMaxError = err.Error()
			} else {
				mStats.ThreadsMax = tmp
			}
			// vm.max_map_count
			if tmp, err := osKit.GetMaxMapCount(); err != nil {
				mStats.MaxMapCountError = err.Error()
			} else {
				mStats.MaxMapCount = tmp
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
