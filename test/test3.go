package main

import (
	"fmt"
)

func main() {
	fmt.Println("c", test()) // c 2
}

func test() (rst int) {
	defer fmt.Println("a", rst) // a 0
	defer func() {
		fmt.Println("b", rst) // b 2
	}()

	rst = 1
	return 2 // 可以拆分为2行代码: rst = 2 && return
}
