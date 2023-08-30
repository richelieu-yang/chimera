package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/statKit"
)

func main() {
	stats := statKit.GetStats()
	fmt.Println(jsonKit.MarshalIndentToString(stats, "", "    "))
}
