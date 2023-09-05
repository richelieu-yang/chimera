package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/compareKit"
)

func main() {
	str := compareKit.Diff("1234", "1234")
	fmt.Println(str)
	fmt.Println(str == "")
}
