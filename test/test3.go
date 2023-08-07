package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func main() {
	fmt.Println(fileKit.IsHidden("/Users/richelieu/go1.21rc3"))
}
