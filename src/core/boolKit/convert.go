package boolKit

import (
	"github.com/spf13/cast"
	"strings"
)

var (
	ToBool = cast.ToBool

	ToBoolE = cast.ToBoolE
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

func ToBoolWithDefault(src interface{}, def bool) bool {
	flag, err := ToBoolE(src)
	if err != nil {
		return def
	}
	return flag
}
