package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
)

func main() {
	m := map[string]interface{}{
		"1": 1,
	}
	fmt.Println(m)

	clear()

	mapKit.Clear(m)
	fmt.Println(m)
}
