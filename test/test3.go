package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"math"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)
}
