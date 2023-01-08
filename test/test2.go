package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	var s []int = nil
	sliceKit.Swap(s, 0, 2)
	fmt.Println(s) // [100 0 1]
}
