package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/interfaceKit"
	"github.com/richelieu42/chimera/src/core/ptrKit"
)

func main() {
	ptr := ptrKit.ToPtr[interface{}](nil)
	fmt.Println(ptr)                     // 0x1400010c5c0
	fmt.Println(ptr == nil)              // false
	fmt.Println(interfaceKit.IsNil(ptr)) // false
}
