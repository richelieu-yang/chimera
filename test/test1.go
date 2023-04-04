package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/interfaceKit"
)

func main() {
	fmt.Println(interfaceKit.IsZeroValue[interface{}](nil))
}
