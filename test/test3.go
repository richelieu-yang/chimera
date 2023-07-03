package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
)

type SampleStruct struct {
	Value string
}

func main() {
	s := sliceKit.Merge[string](nil, []string{})
	fmt.Println(s)        // []
	fmt.Println(len(s))   // 0
	fmt.Println(s != nil) // true
}
