package main

import (
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

type User struct {
	Emails []string `validate:"required"`
}

func main() {
	u := &User{
		Emails: []string{},
	}
	v := validateKit.New()
	if err := v.Struct(u); err != nil {
		panic(err) // panic: Key: 'User.Emails' Error:Field validation for 'Emails' failed on the 'gt' tag
	}
}
