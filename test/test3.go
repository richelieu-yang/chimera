package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/osKit"
)

func main() {
	fmt.Println(osKit.GetThreadCount())
	fmt.Println(osKit.GetProcessCount())
}
