package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func main() {
	fmt.Println(strKit.ContainsIgnoreCase("abcde", "BC"))
}
