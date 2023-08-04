package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"time"
)

func main() {
	path := "cyy.log"

	f, err := fileKit.NewFileInAppendMode(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	logger := logrusKit.NewLogger(logrusKit.WithOutput(f))

	go func() {
		time.Sleep(time.Second * 10)

		if err := fileKit.Truncate(path, 0); err != nil {
			panic(err)
		}
		fmt.Println("------")
	}()

	for i := 0; i < 10000; i++ {
		logger.Info(i)
		time.Sleep(time.Second)
	}
}

//type (
//	MemoryStat struct {
//		// Alloc bytes of allocated heap objects
//		Alloc uint64
//
//		// TotalAlloc cumulative bytes allocated for heap objects
//		TotalAlloc uint64
//
//		// Sys total bytes of memory obtained from the OS
//		Sys uint64
//
//		// NumGC the number of completed GC cycles
//		NumGC uint32
//	}
//)
//
//func main() {
//	var m runtime.MemStats
//	runtime.ReadMemStats(&m)
//
//	logrus.Info(m.Alloc)
//	logrus.Info(m.TotalAlloc)
//	logrus.Info(m.Sys)
//	logrus.Info(m.NumGC)
//
//	logrus.Info("-------")
//
//	memStat, err := memoryKit.GetMemoryStat()
//	if err != nil {
//		panic(err)
//	}
//	logrus.Info(memoryKit.MemoryStatToString(memStat))
//
//	//logrus.Info(memStat.Total)
//	//logrus.Info(memStat.Available)
//	//logrus.Info(memStat.Used)
//	//logrus.Info(memStat.UsedPercent)
//	//logrus.Info(memStat.Free)
//	//
//	//logrus.Info(memStat.Total - memStat.Used - memStat.Available)
//	//
//	////logx.Statf("MEMORY: Alloc=%.1fMi, TotalAlloc=%.1fMi, Sys=%.1fMi, NumGC=%d",
//	////	bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
//}
