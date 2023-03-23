package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/pathKit"
)

func main() {
	path := "/Users/richelieu/Downloads/"
	a, b := pathKit.Split(path)
	fmt.Println(a)
	fmt.Println(b)
}
