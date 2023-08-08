package statKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"runtime"
)

type (
	Stats struct {
		Program *ProgramStats `json:"program"`

		Machine      *MachineStats `json:"machine"`
		MachineError error         `json:"machineError"`
	}

	ProgramStats struct {
		GoroutineCount int

		Alloc      uint64
		TotalAlloc uint64
		Sys        uint64
		NumGC      uint32
		EnableGC   bool
	}

	MachineStats struct {
		// ProcessCount 进程数
		ProcessCount      uint32
		ProcessCountError error

		// ProcessThreadCount 进程数（包括线程数）
		ProcessThreadCount      uint32
		ProcessThreadCountError error
	}
)

func GetStats() (rst *Stats) {
	rst = &Stats{}

	var pStats = &ProgramStats{}
	{
		stats := memoryKit.GetProgramMemoryStats()

		pStats.GoroutineCount = runtime.NumGoroutine()

		pStats.Alloc = stats.Alloc
		pStats.TotalAlloc = stats.TotalAlloc
		pStats.Sys = stats.Sys
		pStats.NumGC = stats.NumGC
		pStats.EnableGC = stats.EnableGC
	}
	rst.Program = pStats

	var mStats *MachineStats
	{
		stats, err := memoryKit.GetMachineMemoryStats()
		if err != nil {
			rst.MachineError = err
		} else {
			mStats = &MachineStats{}

		}
	}
	rst.Machine = mStats

}
