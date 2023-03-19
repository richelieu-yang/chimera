package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/mathKit"
)

func main() {
	fmt.Println(mathKit.Clamp(0, -10, 10))
	fmt.Println(mathKit.Clamp(-42, -10, 10))
	fmt.Println(mathKit.Clamp(42, -10, 10))
}
