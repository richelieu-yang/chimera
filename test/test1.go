package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/mapKit"
)

func main() {
	m := mapKit.Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"a": 10},
		map[string]int{"a": 100},
	)
	fmt.Println(m) // map[a:100 b:2]
}
