package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

func main() {
	test()
}

func test() {
	fmt.Println(errorKit.Simple("Simple"))
	fmt.Println(errorKit.New("New"))
	fmt.Println(errorKit.WithMessage(errorKit.New("New"), "wmq"))
}
