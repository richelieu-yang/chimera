package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func main() {
	test()
}

func test() {
	fmt.Println(funcKit.GetFuncName(1))
}
