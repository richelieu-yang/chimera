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
	fmt.Println(dataSizeKit.ToReadableStringWithIEC(150600))
}
