package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Var("", "omitempty,gte=6"))
	fmt.Println(validateKit.Var("public", "omitempty,gte=6"))
	fmt.Println(validateKit.Var("publi", "omitempty,gte=6"))
}
