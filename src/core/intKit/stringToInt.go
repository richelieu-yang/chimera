package intKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"strconv"
)

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

// ParseStringToInt64 类型转换: string => int64
func ParseStringToInt64(str string) (int64, error) {
	str = strKit.RemoveSpace(str)

	return strconv.ParseInt(str, 10, 64)
}

// ParseStringToInt64WithDefault 类型转换: string => int
func ParseStringToInt64WithDefault(str string, def int64) int64 {
	i, err := ParseStringToInt64(str)
	if err != nil {
		return def
	}
	return i
}
