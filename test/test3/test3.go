package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

func main() {
	fmt.Println(sliceKit.Range(-4, 4))            // [-4 -3 -2 -1]
	fmt.Println(sliceKit.RangeWithStep(-4, 1, 2)) // [-4 -2 0]
}
