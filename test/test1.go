package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	slice := sliceKit.Replace([]int{0, 1, 2, 2, 3, 3, 3}, 3, 9, -1)
	fmt.Println(slice)
}
