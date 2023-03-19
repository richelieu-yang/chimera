package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/samber/lo"
	"unsafe"
)

func main() {

	lo.Reduce()

	var s0 []string = nil
	s1 := sliceKit.Reverse(s0)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(unsafe.Pointer(&s0))
	fmt.Println(unsafe.Pointer(&s1))
}
