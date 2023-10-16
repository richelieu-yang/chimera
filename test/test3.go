package main

import (
	"fmt"
)

func main() {
	fmt.Println(test()) // 1
}

func test() int {
	a := 1

	defer func() {
		a = 2
	}()

	return a
}
