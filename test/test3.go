package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2}
	fmt.Println(s) // [0 1 2]

	sliceKit.ForEach(s, func(item int, index int) {
		s[index] = item + 1
	})
	fmt.Println(s) // [1 2 3]
}
