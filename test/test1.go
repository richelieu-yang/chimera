package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		panic("len(args) == 0")
	}
	path := args[0]
	fmt.Println("path:", path)
	fmt.Println(fileKit.IsHidden(path))
}
