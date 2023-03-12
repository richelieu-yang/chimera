package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
)

func main() {
	path := "/Users/richelieu/START"
	fmt.Println(fileKit.IsHidden(path))
}
