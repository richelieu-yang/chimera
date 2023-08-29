package jsonKit

import "github.com/tidwall/gjson"

// GetStringField
/*
@param path 如果是嵌套的内部字段，可以通过'.'组合
*/
func GetStringField(jsonData []byte, path string) string {
	result := gjson.GetBytes(jsonData, path)
	return result.String()
}

func GetStringFieldFromString(jsonStr, path string) string {
	result := gjson.Get(jsonStr, path)
	return result.String()
}

func GetInt64Field(jsonData []byte, path string) int64 {
	result := gjson.GetBytes(jsonData, path)
	return result.Int()
}

func GetInt64FieldFromString(jsonStr, path string) int64 {
	result := gjson.Get(jsonStr, path)
	return result.Int()
}

func GetBoolField(jsonData []byte, path string) bool {
	result := gjson.GetBytes(jsonData, path)
	return result.Bool()
}

func GetBoolFieldFromString(jsonStr, path string) bool {
	result := gjson.Get(jsonStr, path)
	return result.Bool()
}
