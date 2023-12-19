package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type Bean struct {
	Status string `validate:"eq=active"`
	Age    int    `validate:"required_if=Status active"`
}

func main() {
	err := validateKit.Struct(&Bean{
		Status: "active",
		Age:    0,
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

}
