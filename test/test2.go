package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/pathKit"
)

func main() {
	fmt.Println(pathKit.IsAbs("/root"))
	fmt.Println(pathKit.IsAbs("root"))
}
