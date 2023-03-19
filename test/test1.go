package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"strconv"
)

func main() {
	s := sliceKit.ConvertElementTypeInParallel([]int{0, 1, 2, 3}, func(item int, index int) string {
		return "0x" + strconv.Itoa(item)
	})
	fmt.Println(s) // [0x0 0x1 0x2 0x3]
}
