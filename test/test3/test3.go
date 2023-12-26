package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/ptrKit"
)

type (
	bean struct {
	}
)

func main() {
	var obj interface{} = nil
	fmt.Println(ptrKit.Of(obj)) // 0x1400008e380

	var b *bean = nil
	fmt.Println(ptrKit.Of(b)) // 0x140000980c0

	var obj1 interface{} = b
	fmt.Println(ptrKit.Of(obj1)) // 0x1400008e390
}
