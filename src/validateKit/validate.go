package validateKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

// New
/*
PS:
(1) 默认的tag name: "validate";
(2) Gin使用的tag name: "binding".

@param tagNameArgs 不传参的话，将采用默认的tag name
*/
func New(tagNameArgs ...string) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	// 修改tag name
	tagName := sliceKit.GetFirstItemWithDefault("", tagNameArgs...)
	if strKit.IsNotEmpty(tagName) {
		v.SetTagName(tagName)
	}

	if err := registerDefaultValidation(v); err != nil {
		panic(err)
	}

	return v
}

// registerDefaultValidation 注册默认的验证器(s).
func registerDefaultValidation(v *validator.Validate) error {
	tag := "port"
	err := v.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		field := fl.Field()
		return netKit.IsValidPort(field)
	})
	if err != nil {
		return err
	}

	return nil
}
