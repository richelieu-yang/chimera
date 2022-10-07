package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
)

// Marshal 序列化.
/*
@param obj 可以为nil
*/
func Marshal(obj interface{}) ([]byte, error) {
	return jsoniter.Marshal(obj)
}

// MarshalToString 序列化为字符串.
/*
PS:
(1) 缺陷: 多次序列化相同的map实例(length >= 2)，返回值可能不同，想解决可以使用 MarshalToStringWithJsoniterApi().

@param obj 可以为nil

e.g.
(nil) => ("null", nil)
*/
func MarshalToString(obj interface{}) (string, error) {
	if str, err := jsoniter.MarshalToString(obj); err != nil {
		return "", err
	} else {
		return str, nil
	}
}

// MarshalToStringWithJsoniterApi
/*
@param api 可以为nil，jsoniter.ConfigDefault || jsoniter.ConfigCompatibleWithStandardLibrary || ...

使用场景：多次序列化相同的map实例，希望返回值相同（此时需要传参 jsoniter.ConfigCompatibleWithStandardLibrary）.
*/
func MarshalToStringWithJsoniterApi(api jsoniter.API, obj interface{}) (string, error) {
	if api == nil {
		api = jsoniter.ConfigDefault
	}
	if str, err := api.MarshalToString(obj); err != nil {
		return "", err
	} else {
		return str, nil
	}
}

func MarshalWithIndent(obj interface{}) ([]byte, error) {
	prefix := ""
	// Richelieu: 与golang自带的"encoding/json"（可以用"\t"）不同，indent中不能有非空格的字符
	//indent := "\t"
	indent := "    "

	if data, err := jsoniter.MarshalIndent(obj, prefix, indent); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

// MarshalToStringWithIndent
/*
PS: string(nil) => ""
*/
func MarshalToStringWithIndent(obj interface{}) (string, error) {
	data, err := MarshalWithIndent(obj)
	if err != nil {
		return "", err
	} else {
		return string(data), nil
	}
}
