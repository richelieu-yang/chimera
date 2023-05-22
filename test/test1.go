package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/copyKit"
)

func main() {
	a, err := copyKit.DeepCopy[interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
	fmt.Println(a == nil) // true

	b, err := copyKit.DeepCopy[[]int](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(b == nil) // true

	c, err := copyKit.DeepCopy[map[string]interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(c)
	fmt.Println(c == nil) // true
}
