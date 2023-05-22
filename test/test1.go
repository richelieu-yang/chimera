package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/interfaceKit"
)

func main() {
	var src interface{} = nil
	var src1 []string = nil
	var src2 map[string]interface{} = nil
	type bean struct {
	}
	var src3 *bean = nil

	fmt.Println(interfaceKit.IsNil(src))  // true
	fmt.Println(interfaceKit.IsNil(src1)) // true
	fmt.Println(interfaceKit.IsNil(src2)) // true
	fmt.Println(interfaceKit.IsNil(src3)) // true
}
