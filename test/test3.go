package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func main() {
	fmt.Println(jsonKit.MarshalIndentToString(map[string]interface{}{
		"a": true,
	}, "", "    "))
}
