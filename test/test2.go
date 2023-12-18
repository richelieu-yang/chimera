package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Status string
	Bean   *bean `validate:"required_if=Status active"`
}

type bean struct {
	A int `validate:"gte=10"`
	B bool
}

func main() {
	validate := validator.New()

	user := &User{
		Status: "active111",
		Bean: &bean{
			A: 1,
			B: false,
		},
	}
	err := validate.Struct(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("ok")
}
