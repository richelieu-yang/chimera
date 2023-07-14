package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/runtimeKit"
)

func main() {
	fmt.Println(runtimeKit.GoVersion)

	//fmt.Println(max(0, -1, 1, 2, 3))
	fmt.Println(mathKit.Max(0, -1, 1, 2, 3))
}
