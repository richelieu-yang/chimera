package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"runtime"
)

type (
	Stats struct {
		Program *ProgramStats `json:"program"`

		Machine *MachineStats `json:"machine"`
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

func GetStats() (rst *Stats) {
	rst = &Stats{}

	/* program */
	var pStats = &ProgramStats{}
	rst.Program = pStats
	{
		stats := memoryKit.GetProgramMemoryStats()

		pStats.GoroutineCount = runtime.NumGoroutine()

		pStats.Alloc = dataSizeKit.ToReadableStringWithIEC(stats.Alloc)
		pStats.TotalAlloc = dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc)
		pStats.Sys = dataSizeKit.ToReadableStringWithIEC(stats.Sys)
		pStats.NumGC = stats.NumGC
		pStats.EnableGC = stats.EnableGC
	}

	/* machine */
	var mStats = &MachineStats{}
	rst.Machine = mStats
	{
		count, err := osKit.GetProcessCount()
		if err != nil {
			mStats.ProcessCountError = err
		} else {
			mStats.ProcessCount = count
		}

		count1, err := osKit.GetProcessThreadCount()
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
			mStats.UsedPercent = stats.UsedPercent
			mStats.Free = dataSizeKit.ToReadableStringWithIEC(stats.Free)
		}
	}

	return rst
}
