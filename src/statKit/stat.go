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
		ProcessCount      int   `json:"processCount"`
		ProcessCountError error `json:"processCountError"`

		// ProcessThreadCount 进程数（包括线程数）
		ProcessThreadCount      int   `json:"processThreadCount"`
		ProcessThreadCountError error `json:"processThreadCountError"`

		MemoryStatsError error   `json:"memoryStatsError"`
		Total            uint64  `json:"total"`
		Available        uint64  `json:"available"`
		Used             uint64  `json:"used"`
		UsedPercent      float64 `json:"usedPercent"`
		Free             uint64  `json:"free"`
	}
)

func GetStats() (rst *Stats) {
	rst = &Stats{}

	var pStats = &ProgramStats{}
	{
		stats := memoryKit.GetProgramMemoryStats()

		pStats.GoroutineCount = runtime.NumGoroutine()

		pStats.Alloc = dataSizeKit.ToReadableStringWithIEC(stats.Alloc)
		pStats.TotalAlloc = dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc)
		pStats.Sys = dataSizeKit.ToReadableStringWithIEC(stats.Sys)
		pStats.NumGC = stats.NumGC
		pStats.EnableGC = stats.EnableGC
	}
	rst.Program = pStats

	var mStats = &MachineStats{}
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
			mStats.Total = stats.Total
			mStats.Available = stats.Available
			mStats.Used = stats.Used
			mStats.UsedPercent = stats.UsedPercent
			mStats.Free = stats.Free
		}
	}
	rst.Machine = mStats

	return rst
}
