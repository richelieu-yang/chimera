package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/sonicKit"
)

func main() {
	fmt.Println(sonicKit.MarshalToString(666))
}
