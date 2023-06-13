package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
)

func main() {
	var i uint64 = 9223372036854775807
	fmt.Println(dataSizeKit.ToReadableStringWithIEC(i))
}
