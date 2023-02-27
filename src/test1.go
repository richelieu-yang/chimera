package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s0 := []string{"0", "1", "2"}
	s1 := sliceKit.Copy(s0)

	s1[0] = "3"
	fmt.Println(s0) // [0 1 2]
	fmt.Println(s1) // [3 1 2]
}
