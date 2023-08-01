package main

import (
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

	mem

	//logx.Statf("MEMORY: Alloc=%.1fMi, TotalAlloc=%.1fMi, Sys=%.1fMi, NumGC=%d",
	//	bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

func bToMb(b uint64) float32 {
	return float32(b) / 1024 / 1024
}
