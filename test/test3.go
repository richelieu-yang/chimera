package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type User struct {
	Emails []string `validate:"required"`
}

func main() {
	u := &User{
		Emails: []string{},
	}
	u = nil
	if err := validateKit.Struct(u); err != nil {
		fmt.Println(err.Error())
		panic(err) //
	}
}
