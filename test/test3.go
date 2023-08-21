package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
)

func main() {
	mathutil.Abs()

	rst := mathKit.Exponent(2, 10)
	fmt.Println(rst) // 1024
}
