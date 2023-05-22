package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/ptrKit"
)

func main() {
	type bean struct {
	}
	var b *bean = nil
	fmt.Println(ptrKit.IsPointer(b)) // true（类型为指针，虽然值为nil）

	fmt.Println(ptrKit.IsPointer(interface{}(nil))) // false
}
