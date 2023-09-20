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

// Field 验证字段.
func Field(field interface{}, tag string) error {
	v := New()
	return v.Var(field, tag)
}

// Required 必填，非零值（zero value）
/*
	e.g.
		fmt.Println(validateKit.Required(nil)) 		// Key: '' Error:Field validation for '' failed on the 'required' tag

		fmt.Println(validateKit.Required(""))    	// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required("aaa")) 	// <nil>

		fmt.Println(validateKit.Required(0)) 		// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(1)) 		// <nil>

		fmt.Println(validateKit.Required(false)) 	// Key: '' Error:Field validation for '' failed on the 'required' tag
		fmt.Println(validateKit.Required(true))  	// <nil>
*/
func Required(field interface{}) error {
	return Field(field, "required")
}

// IP
/*
	e.g.
		fmt.Println(validateKit.IP(""))          // Key: '' Error:Field validation for '' failed on the 'ip' tag
		fmt.Println(validateKit.IP("127.0.0.1")) // <nil>
		fmt.Println(validateKit.IP("127.001"))   // Key: '' Error:Field validation for '' failed on the 'ip' tag
*/
func IP(field interface{}) error {
	return Field(field, "ip")
}

func IPv4(field interface{}) error {
	return Field(field, "ipv4")
}

func Email(field interface{}) error {
	return Field(field, "email")
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
	return Field(field, "http_url")
}

// Json 字符串值是否为有效的JSON.
/*
	e.g.
		fmt.Println(validateKit.Json(""))   // Key: '' Error:Field validation for '' failed on the 'json' tag
		fmt.Println(validateKit.Json("[]")) // <nil>
		fmt.Println(validateKit.Json("{}")) // <nil>
		fmt.Println(validateKit.Json("[}")) // Key: '' Error:Field validation for '' failed on the 'json' tag
*/
func Json(field interface{}) error {
	return Field(field, "json")
}

// File 字符串值是否包含有效的文件路径，以及该文件是否存在于计算机上.
/*
	PS: 传参对应的应当是"文件"，是"目录"的话会返回error.

	e.g.
		fmt.Println(validateKit.File("")) // Key: '' Error:Field validation for '' failed on the 'file' tag

		// 目录存在
		fmt.Println(validateKit.File("chimera-lib"))                                         // Key: '' Error:Field validation for '' failed on the 'file' tag
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib")) // Key: '' Error:Field validation for '' failed on the 'file' tag
		// 文件存在
		fmt.Println(validateKit.File("chimera-lib/config.yaml"))                                         // <nil>
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml")) // <nil>
		// 文件不存在
		fmt.Println(validateKit.File("/Users/richelieu/GolandProjects/chimera/chimera-lib/config111.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
		// 无效的文件路径
		fmt.Println(validateKit.File("chimera-lib\\config.yaml")) // Key: '' Error:Field validation for '' failed on the 'file' tag
*/
func File(field interface{}) error {
	return Field(field, "file")
}
