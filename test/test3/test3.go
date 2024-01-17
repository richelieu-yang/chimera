package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
)

func main() {
	var s []int = nil
	fmt.Println(sliceKit.InterceptAfter(s, 0)) // [1]
	fmt.Println(sliceKit.InterceptAfter(s, 1)) // []
}
