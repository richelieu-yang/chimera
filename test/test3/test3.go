package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
	"math"
	"net"
)

func main() {
	net.ParseIP()

	fmt.Println(mathKit.Round(math.NaN(), 2))
	fmt.Println(mathKit.Round(-math.NaN(), 2))
}
