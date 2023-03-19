package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s := sliceKit.Filter([]int{0, 1, 2, 3}, func(item int, index int) bool {
		return item >= 2
	})
	fmt.Println(s) // [2 3]
}
