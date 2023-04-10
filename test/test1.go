package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
)

func main() {
	s := sliceKit.Copy([]int(nil))
	fmt.Println(s)        // []
	fmt.Println(len(s))   // 0
	fmt.Println(s != nil) // true
}
