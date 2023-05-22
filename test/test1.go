package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/copyKit"
)

func main() {
	dest, err := copyKit.DeepCopy[interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(dest == nil) // true

	dest, err = copyKit.DeepCopy[[]int](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(dest == nil) // true

	dest, err = copyKit.DeepCopy[map[string]interface{}](nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(dest == nil) // true

}
