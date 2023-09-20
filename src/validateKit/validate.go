package validateKit

import (
	"github.com/go-playground/validator/v10"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// New
/*
@param tagNameArgs 不传参的话，将采用默认的tag name（"validate"）
*/
func New(tagNameArgs ...string) *validator.Validate {
	v := validator.New(validator.WithRequiredStructEnabled())

	tagName := sliceKit.GetFirstItemWithDefault("", tagNameArgs...)
	if strKit.IsNotEmpty(tagName) {
		v.SetTagName(tagName)
	}

	return v
}

// ValidateField 验证字段.
func ValidateField(field interface{}, tag string) error {
	v := New()
	return v.Var(field, tag)
}

func IPv4(field interface{}) error {
	return ValidateField(field, "ipv4")
}

func Email(field interface{}) error {
	return ValidateField(field, "email")
}

func HttpUrl(field interface{}) error {
	return ValidateField(field, "http_url")
}
