package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/statKit"
)

func main() {
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
}
