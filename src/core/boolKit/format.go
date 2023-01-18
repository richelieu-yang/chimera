package boolKit

import (
	"github.com/spf13/cast"
	"strconv"
)

// BoolToString 类型转换: bool => string
func BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// StringToBool 类型转换: string => bool
/*
PS: 参考了 strconv.ParseBool()、hutool中的BooleanUtil.toBoolean().
*/
func StringToBool(str string) bool {
	switch str {
	case "1", "t", "T", "true", "TRUE", "True", "yes", "y", "ok", "on", "是", "对", "真", "對", "√":
		return true
	default:
		return false
	}
}

func ToBool(obj interface{}) bool {
	switch obj.(type) {
	case string:
		return StringToBool(obj.(string))
	default:
		return cast.ToBool(obj)
	}
}

func ToBoolE(obj interface{}) (bool, error) {
	switch obj.(type) {
	case string:
		return StringToBool(obj.(string)), nil
	default:
		return cast.ToBoolE(obj)
	}
}

func ParseToBoolWithDefault(src interface{}, def bool) bool {
	flag, err := ToBoolE(src)
	if err != nil {
		return def
	}
	return flag
}
