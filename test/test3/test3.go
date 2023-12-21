package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type (
	bean struct {
		Port int `validate:"nefield=A.Port"`

		A *A
	}
	A struct {
		Port int
	}
)

func main() {
	{
		b := &bean{
			Port: 0,
			A:    &A{Port: 0},
		}
		fmt.Println(validateKit.Struct(b)) // Key: 'bean.Port' Error:Field validation for 'Port' failed on the 'nefield' tag
	}
	{
		b := &bean{
			Port: 0,
			A:    nil,
		}
		fmt.Println(validateKit.Struct(b)) // <nil>
	}
}
