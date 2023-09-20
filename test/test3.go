package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Json(""))   // Key: '' Error:Field validation for '' failed on the 'json' tag
	fmt.Println(validateKit.Json("[]")) // <nil>
	fmt.Println(validateKit.Json("{}")) // <nil>
	fmt.Println(validateKit.Json("[}")) // Key: '' Error:Field validation for '' failed on the 'json' tag
}
