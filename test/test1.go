package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/cmdKit"
)

func main() {
	path := "/root/a b/c  d/java"
	path = cmdKit.PolyfillCommandPath(path)
	fmt.Println(path)
}
