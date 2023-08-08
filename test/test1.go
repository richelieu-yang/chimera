package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/statKit"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: true,
	})

	stats := statKit.GetStats()
	fmt.Println(jsonKit.MarshalIndentToString(stats, "", "    "))

	//fmt.Println("Alloc:", dataSizeKit.ToReadableStringWithIEC(stats.Alloc))
	//fmt.Println("TotalAlloc:", dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc))
	//fmt.Println("Sys:", dataSizeKit.ToReadableStringWithIEC(stats.Sys))
	//fmt.Println("NumGC:", stats.NumGC)
	//
	//fmt.Println(stats.EnableGC)
	//fmt.Println(stats.DebugGC)
	//
	//fmt.Println(stats)

	s, err := process.Pids()
	if err != nil {
		panic(err)
	}
	fmt.Println(len(s))

	//process.PidExists()

}
