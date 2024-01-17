package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	_ "github.com/richelieu-yang/chimera/v2/src/log/logrusInitKit"
)

func main() {
	s := []int{1}
	fmt.Println(sliceKit.InterceptBefore(s, 0)) // []
	fmt.Println(sliceKit.InterceptBefore(s, 1)) // [1]
}
