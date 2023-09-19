package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	path := "/Users/richelieu/Downloads/111.txt"
	fmt.Println(fileKit.IsReadable(path))
	fmt.Println(fileKit.IsWritable(path))
}
