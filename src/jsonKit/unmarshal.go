package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/strKit"
)

// Unmarshal
/**
@param data 必须满足条件: len(data) > 0
@param ptr 	[不能为nil] 必须是指针类型（pointer）
*/
func Unmarshal(data []byte, ptr interface{}) error {
	if len(data) == 0 {
		if data == nil {
			return errorKit.Simple("data == nil")
		}
		return errorKit.Simple("len(data) == 0")
	}
	if err := ptrKit.AssertPointer(ptr, "ptr"); err != nil {
		return err
	}

	return jsoniter.Unmarshal(data, ptr)
}

// UnmarshalToMap
/*
@param data 必须满足条件: len(data) > 0
*/
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
	if err := ptrKit.AssertPointer(ptr, "ptr"); err != nil {
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
