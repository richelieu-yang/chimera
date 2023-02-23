package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/mapKit"
)

func main() {
	m := map[string]interface{}{
		"a": 0,
		"b": 1,
		"c": 2,
	}

	fmt.Println(mapKit.Remove(m, "b"))
	fmt.Println(m)
}
