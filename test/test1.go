package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []int{0, 1, 2}
	s1 := sliceKit.Intercept(s, len(s), len(s))
	fmt.Println(s1)        // []
	fmt.Println(s1 != nil) // true
}
