package intKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

func ToInt(obj interface{}) int {
	return cast.ToInt(obj)
}

func ToIntE(obj interface{}) (int, error) {
	return cast.ToIntE(obj)
}

func ToInt32(obj interface{}) int32 {
	return cast.ToInt32(obj)
}

func ToInt32E(obj interface{}) (int32, error) {
	return cast.ToInt32E(obj)
}

func ToInt64(obj interface{}) int64 {
	return cast.ToInt64(obj)
}

func ToInt64E(obj interface{}) (int64, error) {
	return cast.ToInt64E(obj)
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
