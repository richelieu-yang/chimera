package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/reflectKit"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"John", 30}
	fmt.Println(p)

	if err := reflectKit.SetField(&p, "name", "Tom"); err != nil {
		panic(err)
	}
	fmt.Println(p)
}
