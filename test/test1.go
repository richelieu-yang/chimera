package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/mapKit"
)

func main() {
	m := mapKit.Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
		map[string]int{"b": 5},
	)
	fmt.Println(m) // map[a:1 b:5 c:4]
}
