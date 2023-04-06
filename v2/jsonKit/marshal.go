package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
)

// Marshal 序列化.
/*
@param obj 可以为nil || ""

e.g.
(nil) 	=> []byte("null"), nil
("") 	=> []byte("\"\""), nil
*/
func Marshal(obj interface{}) ([]byte, error) {
	return jsoniter.Marshal(obj)
}

// MarshalToString 序列化为字符串.
/*
PS:
(1) 缺陷: 多次序列化相同的map实例(length >= 2)，返回值可能不同，想解决可以使用 MarshalToStringWithJsoniterApi().

@param obj 可以为nil || ""

e.g.
(nil) 	=> "null", nil
("") 	=> "\"\"", nil
*/
func MarshalToString(obj interface{}) (string, error) {
	return jsoniter.MarshalToString(obj)
}

// MarshalWithJsoniterApi 可以自定义api
func MarshalWithJsoniterApi(api jsoniter.API, obj interface{}) ([]byte, error) {
	if api == nil {
		api = jsoniter.ConfigDefault
	}
	return api.Marshal(obj)
}

// MarshalToStringWithJsoniterApi 可以自定义api
/*
@param api 可以为nil，jsoniter.ConfigDefault || jsoniter.ConfigCompatibleWithStandardLibrary || ...

e.g.
如果希望多次序列化同一map实例，返回的json字符串一直，传参api可以为 jsoniter.ConfigCompatibleWithStandardLibrary.
*/
func MarshalToStringWithJsoniterApi(api jsoniter.API, obj interface{}) (string, error) {
	if api == nil {
		api = jsoniter.ConfigDefault
	}
	return api.MarshalToString(obj)
}

func MarshalWithIndent(obj interface{}) ([]byte, error) {
	prefix := ""
	// Richelieu: 与golang自带的"encoding/json"（可以用"\t"）不同，indent中不能有非空格的字符
	//indent := "\t"
	indent := "    "

	return jsoniter.MarshalIndent(obj, prefix, indent)
}

func MarshalToStringWithIndent(obj interface{}) (string, error) {
	data, err := MarshalWithIndent(obj)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
