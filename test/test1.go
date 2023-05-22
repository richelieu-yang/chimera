package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/copyKit"
)

func main() {
	src := []string{"0", "1", "2"}
	dest := make([]string, len(src))
	if err := copyKit.Copy(&dest, src); err != nil {
		panic(err)
	}

	src[0] = "a"
	fmt.Println(src)  // [a 1 2]
	fmt.Println(dest) // [0 1 2]
}
