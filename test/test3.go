package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"reflect"
)

type config struct {
	A interface{} `json:"a" validate:"port"`
}

func main() {
	v := validator.New(validator.WithRequiredStructEnabled())
	err := v.RegisterValidation("port", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		switch field.Kind() {
		case reflect.String:
			return netKit.IsStringValidPort(field.String())
		default:
			if field.CanInt() {
				return netKit.IsValidPort(field.Int())
			} else if field.CanUint() {
				return netKit.IsUint64ValidPort(field.Uint())
			}
			return false
		}
	})
	if err != nil {
		panic(err)
	}

	c := &config{
		A: true,
	}
	if err := validateKit.Struct(c); err != nil {
		//if err := v.Struct(c); err != nil {
		panic(err)
	}
}
