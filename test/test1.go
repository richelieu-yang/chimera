package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/mapKit"
)

func main() {
	m := map[string]interface{}{
		"a": 1,
		"b": 2,
	}
	fmt.Println(m)

	fmt.Println(mapKit.Remove(m, "b"))
	fmt.Println(m)

	fmt.Println(mapKit.Remove(m, "b"))
	fmt.Println(m)
}
