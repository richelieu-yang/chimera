package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"unsafe"
)

func main() {
	s0 := []string{"0", "1", "2", "3"}
	s1 := sliceKit.Shuffle(s0)

	fmt.Println(s0)
	fmt.Println(s1)
	fmt.Println(unsafe.Pointer(&s0))
	fmt.Println(unsafe.Pointer(&s1))
}
