package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
)

func main() {
	stat, err := runtimeKit.GetDiskStat()
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.String())
}
