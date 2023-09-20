package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.IP(""))          // Key: '' Error:Field validation for '' failed on the 'ip' tag
	fmt.Println(validateKit.IP("127.0.0.1")) // <nil>
	fmt.Println(validateKit.IP("127.001"))   // Key: '' Error:Field validation for '' failed on the 'ip' tag
}
