package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
)

func main() {
	fmt.Println(dataSizeKit.ParseString("42MB"))
	fmt.Println(dataSizeKit.ParseString("42 MB"))
	fmt.Println(dataSizeKit.ParseString("42mib"))
	fmt.Println(dataSizeKit.ParseString("42 mib"))
}
