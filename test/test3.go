package main

import (
	"fmt"
	"github.com/google/wire"
)

func main() {
	wire.NewSet()

	fmt.Println("c", test()) // c 2
}

func test() (rst int) {
	defer fmt.Println("a", rst) // a 0
	defer func() {
		fmt.Println("b", rst) // b 2
	}()

	rst = 1
	rst = 2
	return
}
