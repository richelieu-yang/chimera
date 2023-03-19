package intKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

func ParseToInt32(obj interface{}) (int32, error) {
	switch obj.(type) {
	case string:
		return ParseStringToInt32(obj.(string))
	default:
		return cast.ToInt32E(obj)
	}
}

func ParseToInt32WithDefault(obj interface{}, def int32) int32 {
	i, err := ParseToInt32(obj)
	if err != nil {
		return def
	}
	return i
}

// ParseStringToInt32 类型转换: string => int32
func ParseStringToInt32(str string) (int32, error) {
	str = strKit.RemoveSpace(str)

	i, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// ParseStringToInt32WithDefault 类型转换: string => int
func ParseStringToInt32WithDefault(str string, def int32) int32 {
	i, err := ParseStringToInt32(str)
	if err != nil {
		return def
	}
	return i
}
