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

func requiredIf(fl FieldLevel) bool {
	params := parseOneOfParam2(fl.Param())
	if len(params)%2 != 0 {
		panic(fmt.Sprintf("Bad param number for required_if %s", fl.FieldName()))
	}
	for i := 0; i < len(params); i += 2 {
		if !requireCheckFieldValue(fl, params[i], params[i+1], false) {
			return true
		}
	}
	return hasValue(fl)
}
