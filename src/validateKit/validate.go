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

// Required
/*
	e.g.
		fmt.Println(validateKit.Required(""))    // Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(nil))   // Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required("aaa")) // <nil>
		fmt.Println(validateKit.Required(1))     // <nil>
*/
func Required(field interface{}) error {
	return ValidateField(field, "required")
}

// IP
/*
	e.g.
		fmt.Println(validateKit.IP(""))          // Key: '' Error:Field validation for '' failed on the 'ip' tag
		fmt.Println(validateKit.IP("127.0.0.1")) // <nil>
		fmt.Println(validateKit.IP("127.001"))   // Key: '' Error:Field validation for '' failed on the 'ip' tag
*/
func IP(field interface{}) error {
	return ValidateField(field, "ip")
}

func IPv4(field interface{}) error {
	return ValidateField(field, "ipv4")
}

func Email(field interface{}) error {
	return ValidateField(field, "email")
}

// HttpUrl
/*
	PS: 要以 "http://" 或 "https://" 开头.

	e.g.
		fmt.Println(validateKit.HttpUrl(""))                                           // Key: '' Error:Field validation for '' failed on the 'http_url' tag
		fmt.Println(validateKit.HttpUrl("https://github.com/go-playground/validator")) // <nil>
		fmt.Println(validateKit.HttpUrl("http://github.com/go-playground/validator"))  // <nil>
		fmt.Println(validateKit.HttpUrl("ftp://github.com/go-playground/validator"))   // Key: '' Error:Field validation for '' failed on the 'http_url' tag
*/
func HttpUrl(field interface{}) error {
	return ValidateField(field, "http_url")
}

// Json
/*
	e.g.
		fmt.Println(validateKit.Json(""))   // Key: '' Error:Field validation for '' failed on the 'json' tag
		fmt.Println(validateKit.Json("[]")) // <nil>
		fmt.Println(validateKit.Json("{}")) // <nil>
		fmt.Println(validateKit.Json("[}")) // Key: '' Error:Field validation for '' failed on the 'json' tag
*/
func Json(field interface{}) error {
	return ValidateField(field, "json")
}
