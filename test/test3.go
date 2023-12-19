package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type bean struct {
	// A 可以为: "" || 长度为[2, 3]的字符串
	A string `validate:"omitempty,min=2,max=3"`
}

func main() {
	{
		b := &bean{
			A: "",
		}

		err := validateKit.Struct(b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok") // ok
		}
	}
	{
		b := &bean{
			A: "a",
		}

		err := validateKit.Struct(b)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok") // Key: 'bean.A' Error:Field validation for 'A' failed on the 'min' tag
		}
	}
}
