package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2, 3}
	s1, item, ok := sliceKit.RemoveByIndex(s, 2)

	fmt.Println(s)    // [0 1 2 3]
	fmt.Println(s1)   // [0 1 3]
	fmt.Println(item) // 2
	fmt.Println(ok)   // true
}
