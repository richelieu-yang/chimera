package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	fmt.Println(fileKit.GetFileMode("111"))
	fmt.Println(fileKit.GetFileMode("Makefile"))
}
