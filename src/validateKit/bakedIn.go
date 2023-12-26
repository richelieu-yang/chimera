package validateKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

var bakedInValidators = map[string]validator.Func{
	"port": isPort,
}

func isPort(fl validator.FieldLevel) bool {
	field := fl.Field()
	return netKit.IsValidPort(field)
}

func registerBakedInValidation(v *validator.Validate) (err error) {
	for tag, fn := range bakedInValidators {
		err = v.RegisterValidation(tag, fn)
		if err != nil {
			return
		}
	}
	return
}
