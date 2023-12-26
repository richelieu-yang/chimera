package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/ptrKit"
)

func main() {
	var ptr *string = nil
	fmt.Println(ptrKit.Unwrap(ptr)) // panic: runtime error: invalid memory address or nil pointer dereference
}
