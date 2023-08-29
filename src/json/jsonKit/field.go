package jsonKit

import "github.com/tidwall/gjson"

// GetStringField
/*
@param path 如果是嵌套的内部字段，可以通过'.'组合
*/
func GetStringField(json []byte, path string) string {
	result := gjson.GetBytes(json, path)
	return result.String()
}

func GetStringFieldFromString(json, path string) string {
	result := gjson.Get(json, path)
	return result.String()
}
