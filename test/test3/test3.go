package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
)

func main() {
	fmt.Println(mathKit.Factorial(0)) // 1
	fmt.Println(mathKit.Factorial(1)) // 1
	fmt.Println(mathKit.Factorial(5)) // 120（=1*2*3*4*5）
}
