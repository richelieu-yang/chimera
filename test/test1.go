package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
)

func main() {
	src := []string{"0", "1", "2"}
	dest := sliceKit.Copy(src)

	src[0] = "a"
	fmt.Println(src)  // [a 1 2]
	fmt.Println(dest) // [0 1 2]
}
