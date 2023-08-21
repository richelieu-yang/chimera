package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/mathKit"
)

func main() {
	result0 := mathKit.TruncRound(1234.124, 0)
	result1 := mathKit.TruncRound(1234.124, -1)
	result2 := mathKit.TruncRound(1234.124, -2)

	result3 := mathKit.TruncRound(100.125, 2)
	result4 := mathKit.TruncRound(100.125, 3)

	fmt.Println(result0)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}
