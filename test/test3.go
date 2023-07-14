package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
)

func main() {
	fmt.Println(runtimeKit.GoVersion)
	m := map[string]interface{}{
		"1": 1,
		"0": 0,
	}
	fmt.Println(m)

	mapKit.Clear(m)
	fmt.Println(m)
}
