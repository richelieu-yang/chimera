package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"math"
)

func main() {
	fmt.Println(dataSizeKit.ToReadableStringWithIEC(math.MaxInt32))
}
