package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Var("", "omitempty,hostname_port"))                // <nil>
	fmt.Println(validateKit.Var("127.0.0.1:80", "omitempty,hostname_port"))    // <nil>
	fmt.Println(validateKit.Var("127.0.0.1:65535", "omitempty,hostname_port")) // <nil>

	fmt.Println(validateKit.Var("127.0.0.1:65536", "omitempty,hostname_port")) // Key: '' Error:Field validation for '' failed on the 'hostname_port' tag
}
