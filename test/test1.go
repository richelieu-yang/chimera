package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s0 := []int{0, 1, 2, 3}
	s1 := sliceKit.Filter(s0, func(item int, index int) bool {
		return item >= 2
	})
	fmt.Println(s0) // [0 1 2 3]
	fmt.Println(s1) // [2 3]
}
