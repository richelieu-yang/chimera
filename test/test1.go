package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/funcKit"
)

func main() {
	fmt.Println(funcKit.GetFuncInfo(0))
	Print()
}

func Print() {
	fmt.Println(funcKit.GetFuncInfo(0))
	test()
}

func test() {
	fmt.Println(funcKit.GetFuncInfo(0))
}
