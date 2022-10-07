package jsonKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/pointerKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	jsoniter "github.com/json-iterator/go"
)

// Unmarshal ！！！：要注意传参data长度为0的情况，会报错
/**
@param ptr 	类型只能为指针（pointer），且不能为nil
*/
func Unmarshal(data []byte, ptr interface{}) error {
	if len(data) == 0 {
		if data == nil {
			return errorKit.Simple("data == nil")
		}
		return errorKit.Simple("len(data) == 0")
	}
	if err := pointerKit.AssertPointer(ptr, "ptr"); err != nil {
		return err
	}

	return jsoniter.Unmarshal(data, ptr)
}

func UnmarshalToMap(data []byte) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// UnmarshalFromString
/*
@param str	!!!: 不能为空字符串("")，否则会报错
@param obj 	只能为指针（pointer），且不能为nil
*/
func UnmarshalFromString(str string, ptr interface{}) error {
	if err := strKit.AssertNotEmpty(str, "str"); err != nil {
		return err
	}
	if err := pointerKit.AssertPointer(ptr, "ptr"); err != nil {
		return err
	}

	return jsoniter.UnmarshalFromString(str, ptr)
}

func UnmarshalFromStringToMap(str string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	err := UnmarshalFromString(str, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}
