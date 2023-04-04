package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/ptrKit"
)

func main() {
	ptr := ptrKit.ToPtr("hello world")
	fmt.Println(ptr) // 0x140000105e0
}
