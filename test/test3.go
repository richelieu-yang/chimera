package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/crypto/unicodeKit"
)

func main() {
	fmt.Println(unicodeKit.Encode("测试")) // "\u6d4b\u8bd5"
}
