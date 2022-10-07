package boolKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

func ParseToBool(src interface{}) (bool, error) {
	switch src.(type) {
	case string:
		return ParseStringToBool(src.(string))
	default:
		if src == nil {
			return false, errorKit.Simple("src == nil")
		}
		if flag, err := cast.ToBoolE(src); err != nil {
			return false, err
		} else {
			return flag, nil
		}
	}
}

func ParseToBoolWithDefault(src interface{}, def bool) bool {
	flag, err := ParseToBool(src)
	if err != nil {
		return def
	}
	return flag
}

// ParseStringToBool 类型转换: string => bool
func ParseStringToBool(str string) (bool, error) {
	ori := str
	str = strKit.RemoveSpace(str)

	if strKit.IsEmpty(str) {
		return false, errorKit.Simple("str(\"%s\") is invalid", ori)
	}

	return strconv.ParseBool(str)
	//switch strings.ToLower(str) {
	//case "true", "yes", "y", "t", "ok", "1", "on", "是", "对", "真", "對", "√":
	//	return true, nil
	//default:
	//	return false, nil
	//}
}

// ParseStringToBoolWithDefault 类型转换: string => bool
/*
参考：strconv.ParseBool()、hutool中的BooleanUtil.toBoolean().
*/
func ParseStringToBoolWithDefault(str string, def bool) bool {
	flag, err := ParseStringToBool(str)
	if err != nil {
		return def
	}
	return flag
}
