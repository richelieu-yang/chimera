package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

func main() {
	str := sliceKit.Join([]string{"1"}, ";")
	fmt.Printf(str)
}
