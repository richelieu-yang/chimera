package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
	"math"
)

func main() {
	fmt.Println(mathKit.Round(math.NaN(), 2))
	fmt.Println(mathKit.Round(-math.NaN(), 2))
}
