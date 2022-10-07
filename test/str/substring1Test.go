package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// main
/**
 * 测试：strKit.Substring()
 */
func main() {
	str := "abc"

	fmt.Println(strKit.Substring(str, 0, 100))

}
