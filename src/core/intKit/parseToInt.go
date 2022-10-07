package intKit

import (
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

func ParseToInt(obj interface{}) (int, error) {
	switch obj.(type) {
	case string:
		return ParseStringToInt(obj.(string))
	default:
		return cast.ToIntE(obj)
	}
}

func ParseToIntWithDefault(obj interface{}, def int) int {
	i, err := ParseToInt(obj)
	if err != nil {
		return def
	}
	return i
}

// ParseStringToInt 类型转换: string => int
func ParseStringToInt(str string) (int, error) {
	str = strKit.RemoveSpace(str)

	return strconv.Atoi(str)
}

// ParseStringToIntWithDefault 类型转换: string => int
func ParseStringToIntWithDefault(str string, def int) int {
	i, err := ParseStringToInt(str)
	if err != nil {
		return def
	}
	return i
}
