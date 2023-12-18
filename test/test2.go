package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Enabled bool
	Reason  string `validate:"excluded_if=Enabled true"`
}

func main() {
	validate := validator.New()

	err := validate.Struct(&User{
		Enabled: true,
		Reason:  "111",
	})
	if err != nil {
		fmt.Println(err) // Key: 'User.Reason' Error:Field validation for 'Reason' failed on the 'excluded_if' tag
	} else {
		fmt.Println("ok")
	}

	err = validate.Struct(&User{
		Enabled: true,
		Reason:  "",
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok") // ok
	}
}
