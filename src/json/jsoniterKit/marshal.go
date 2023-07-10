package jsoniterKit

import "github.com/richelieu-yang/chimera/v2/src/core/strKit"

// Marshal 序列化（可选配置: api && indent）.
/*
Description: 建议使用sonicKit.

@param obj 可以为nil || ""

e.g.
	(nil) 	=> []byte("null"), nil
	("") 	=> []byte("\"\""), nil
*/
func Marshal(obj interface{}, options ...JsonOption) ([]byte, error) {
	opts := loadOptions(options...)

	if strKit.IsAllEmpty(opts.prefix, opts.indent) {
		return opts.api.Marshal(obj)
	}
	return opts.api.MarshalIndent(obj, opts.prefix, opts.indent)
}

// MarshalToString 序列化为json字符串（可选配置: api && indent）.
/*
Description: 建议使用sonicKit.

@param obj 可以为nil || ""

e.g.
	(nil) 	=> "null", nil
	("") 	=> "\"\"", nil

e.g.1
	m := map[string]interface{}{
		"a": "0",
		"b": "1",
	}

	str, _ := jsonKit.MarshalToString(m)
	fmt.Println(str)
	// {"a":"0","b":"1"}

	str, _ = jsonKit.MarshalToString(m, jsonKit.WithIndent("    "))
	fmt.Println(str)
	//{
	//	"a": "0",
	//	"b": "1"
	//}
*/
func MarshalToString(obj interface{}, options ...JsonOption) (string, error) {
	opts := loadOptions(options...)

	if strKit.IsAllEmpty(opts.prefix, opts.indent) {
		return opts.api.MarshalToString(obj)
	}
	data, err := opts.api.MarshalIndent(obj, opts.prefix, opts.indent)
	return string(data), err
}
