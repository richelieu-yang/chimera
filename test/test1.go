package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/funcKit"
)

func main() {
	fmt.Println(funcKit.AddFuncInfoToString("1", 0))
	Print()
}

func Print() {
	fmt.Println(funcKit.AddFuncInfoToString("2", 0))
	test()
}

func test() {
	fmt.Println(funcKit.AddFuncInfoToString("3", 0))
}
