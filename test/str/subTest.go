package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// main
/**
 * 测试：strKit.Sub()
 */
func main() {
	str := "abcd"

	fmt.Println(strKit.Sub(str, 0, 1))  // a
	fmt.Println(strKit.Sub(str, 0, 2))  // ab
	fmt.Println(strKit.Sub(str, 2, 0))  // ab
	fmt.Println(strKit.Sub(str, 2, -1)) // c
	fmt.Println(strKit.Sub(str, 2, 2))  // ""
}
