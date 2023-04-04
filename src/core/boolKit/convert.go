package boolKit

import (
	"github.com/spf13/cast"
	"strings"
)

// StringToBool 类型转换: string => bool
/*
PS: 参考了 strconv.ParseBool()、Java hutool中的BooleanUtil.toBoolean().
*/
func StringToBool(str string) bool {
	switch strings.ToLower(str) {
	case "1", "t", "true", "yes", "y", "ok", "on", "是", "对", "對", "真", "√":
		return true
	default:
		return false
	}
}

// ToBool
/*
e.g.
(nil) => false
*/
func ToBool(obj interface{}) bool {
	switch obj.(type) {
	case string:
		return StringToBool(obj.(string))
	default:
		return cast.ToBool(obj)
	}
}

// ToBoolE
/*
e.g.
(nil) => false, nil
*/
func ToBoolE(obj interface{}) (bool, error) {
	switch obj.(type) {
	case string:
		return StringToBool(obj.(string)), nil
	default:
		return cast.ToBoolE(obj)
	}
}

func ToBoolWithDefault(src interface{}, def bool) bool {
	flag, err := ToBoolE(src)
	if err != nil {
		return def
	}
	return flag
}
