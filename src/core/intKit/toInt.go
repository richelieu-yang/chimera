package intKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/spf13/cast"
	"strconv"
)

// ToInt
/*
e.g.
(nil) => 0
*/
func ToInt(obj interface{}) int {
	return cast.ToInt(obj)
}

// ToIntE
/*
e.g.
(nil) 	=> 0 <nil>
(false)	=> 0 <nil>
(true)	=> 1 <nil>
("")	=> 0 unable to cast "" of type string to int64
*/
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

// StringToInt 类型转换: string => int
func StringToInt(str string) (int, error) {
	str = strKit.RemoveSpace(str)
	return strconv.Atoi(str)
}

// StringToIntWithDefault 类型转换: string => int
func StringToIntWithDefault(str string, def int) int {
	i, err := StringToInt(str)
	if err != nil {
		return def
	}
	return i
}
