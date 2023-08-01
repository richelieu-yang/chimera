package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/sirupsen/logrus"
	"runtime"
)

type (
	MemoryStat struct {
		// Alloc bytes of allocated heap objects
		Alloc uint64

		// TotalAlloc cumulative bytes allocated for heap objects
		TotalAlloc uint64

		// Sys total bytes of memory obtained from the OS
		Sys uint64

		// NumGC the number of completed GC cycles
		NumGC uint32
	}
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	logrus.Info(m.Alloc)
	logrus.Info(m.TotalAlloc)
	logrus.Info(m.Sys)
	logrus.Info(m.NumGC)

	logrus.Info("-------")

	memStat, err := memoryKit.GetMemoryStat()
	if err != nil {
		panic(err)
	}
	logrus.Info(memoryKit.MemoryStatToString(memStat))

	//logrus.Info(memStat.Total)
	//logrus.Info(memStat.Available)
	//logrus.Info(memStat.Used)
	//logrus.Info(memStat.UsedPercent)
	//logrus.Info(memStat.Free)
	//
	//logrus.Info(memStat.Total - memStat.Used - memStat.Available)
	//
	////logx.Statf("MEMORY: Alloc=%.1fMi, TotalAlloc=%.1fMi, Sys=%.1fMi, NumGC=%d",
	////	bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) float32 {
	return float32(b) / 1024 / 1024
}
