package intKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

func ParseToInt64(src interface{}) (int64, error) {
	switch src.(type) {
	case string:
		return ParseStringToInt64(src.(string))
	default:
		return cast.ToInt64E(src)
	}
}

func ParseToInt64WithDefault(src interface{}, def int64) int64 {
	i, err := ParseToInt64(src)
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
