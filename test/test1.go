package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	m := sliceKit.Group[int, int](s, func(i int) int {
		return i % 3
	})
	fmt.Println(s) // [0 1 2 3 4 5]
	fmt.Println(m) // map[0:[0 3] 1:[1 4] 2:[2 5]]
}
