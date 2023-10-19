package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
)

func main() {
	extname := ".JSON"
	fmt.Println(viperKit.PolyfillContentType(extname)) // "json"
}
