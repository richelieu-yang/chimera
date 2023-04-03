package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	var s []int = []int{0, 1, 2, 3}
	s1 := sliceKit.Intercept(s, 1, 0)

	fmt.Println(s1)
	fmt.Println(s1 != nil)

	//s1 := sliceKit.Intercept(s, len(s), len(s))
	//fmt.Println(s1)        // []
	//fmt.Println(s1 != nil) // true
}
