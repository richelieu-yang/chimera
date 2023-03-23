package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	fmt.Println(sliceKit.IsSorted[int](nil))
	fmt.Println(sliceKit.IsSorted([]string{}))
	fmt.Println(sliceKit.IsSorted([]string{"b"}))
	fmt.Println(sliceKit.IsSorted([]string{"b", "a"}))
	fmt.Println(sliceKit.IsSorted([]int{0, 1, 9, 100}))
}
