package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/ptrKit"
)

func main() {
	str := "测试test"
	value := ptrKit.FromPtr(&str)
	fmt.Println(value) // "测试test"
	value = ptrKit.FromPtr[string](nil)
	fmt.Println(value) // ""
}
