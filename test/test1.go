package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s0 := []int{0, 1, 2, 3, 4, 5}
	s1 := sliceKit.UniqBy[int, int](s0, func(i int) int {
		return i % 3
	})
	fmt.Println(s0)
	fmt.Println(s1)
}
