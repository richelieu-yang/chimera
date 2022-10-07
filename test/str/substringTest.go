package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// main
/**
 * 测试：strKit.SubAfter()
 */
func main() {
	str := "abc"

	fmt.Println(strKit.SubAfter(str, 0))  // abc
	fmt.Println(strKit.SubAfter(str, 1))  // bc
	fmt.Println(strKit.SubAfter(str, 2))  // c
	fmt.Println(strKit.SubAfter(str, 3))  // ""
	fmt.Println(strKit.SubAfter(str, 4))  // panic: runtime error: slice bounds out of range [4:3]
	fmt.Println(strKit.SubAfter(str, -1)) // panic: runtime error: slice bounds out of range [-1:]
}
