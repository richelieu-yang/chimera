package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	err := validateKit.Var("allow-from 111", "omitempty,lowercase,oneof=deny sameorigin|startswith=allow-from ")
	fmt.Println(err)
}
