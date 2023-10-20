package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Field(0, "eq=1|eq=2")) // Key: '' Error:Field validation for '' failed on the 'eq=1|eq=2' tag
	fmt.Println(validateKit.Field(1, "eq=1|eq=2")) // <nil>
}
