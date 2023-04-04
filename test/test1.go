package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := sliceKit.Drop(s, 2)
	fmt.Println(s)  // [0 1 2 3 4 5]
	fmt.Println(s1) // [2 3 4 5]

	s1[0] = 9
	fmt.Println(s)  // [0 1 2 3 4 5]
	fmt.Println(s1) // [9 3 4 5]
}
