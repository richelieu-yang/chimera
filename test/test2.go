package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
)

func main() {
	path := "/a/v/c"

	fmt.Println(fileKit.Exist(path))
	fmt.Println(fileKit.Delete(path))
}
