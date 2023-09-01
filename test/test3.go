package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	fmt.Println(dataSizeKit.ToReadableStringWithIEC(9223372036854775807))

	//fmt.Println(math.MinInt64)
	//fmt.Println(math.MaxInt64)
}
