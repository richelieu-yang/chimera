package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"reflect"
)

type config struct {
	A interface{} `json:"a" validate:"port"`
}

func main() {
	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.RegisterValidation("port", func(fl validator.FieldLevel) bool {
		var field reflect.Value = fl.Field()
		return netKit.IsValidPort(field)
	})
	if err != nil {
		panic(err)
	}

	c := &config{
		A: 0,
	}
	//if err := validateKit.Struct(c); err != nil {
	if err := v.Struct(c); err != nil {
		panic(err)
	}
}
