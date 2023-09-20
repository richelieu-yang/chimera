package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Required(nil)) // Key: '' Error:Field validation for '' failed on the 'required' tag

	fmt.Println(validateKit.Required(""))    // Key: '' Error:Field validation for '' failed on the 'required' tag
	fmt.Println(validateKit.Required("aaa")) // <nil>

	fmt.Println(validateKit.Required(0)) // Key: '' Error:Field validation for '' failed on the 'required' tag
	fmt.Println(validateKit.Required(1)) // <nil>

	fmt.Println(validateKit.Required(false)) // Key: '' Error:Field validation for '' failed on the 'required' tag
	fmt.Println(validateKit.Required(true))  // <nil>
}
