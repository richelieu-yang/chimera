package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/diskKit"
)

func main() {
	stat, err := diskKit.GetDiskStat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.String())
	fmt.Println(stat.GetFreePercent())
	fmt.Println(stat.GetUsedPercent())

	fmt.Println(stat.Used + stat.Free - stat.Total)
}
