package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s0 := []int{1, 2, 3, 4}
	sum := sliceKit.Reduce[int, int](s0, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	fmt.Println(s0)  // [1 2 3 4]
	fmt.Println(sum) // 10
}
