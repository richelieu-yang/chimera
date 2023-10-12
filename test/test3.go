package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
		"b": 1,
	}
	fmt.Println(mapKit.Obtain(m, "a")) // 0 true
	fmt.Println(mapKit.Obtain(m, "c")) // <nil> false
}
