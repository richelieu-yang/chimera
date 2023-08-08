package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/disk"
)

func main() {
	parts, err := disk.Partitions(true)
	if err != nil {
		panic(err)
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		panic(err)
	}
	fmt.Println(diskInfo.UsedPercent)

}
