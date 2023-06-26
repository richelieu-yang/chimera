package intKit

import (
	"github.com/spf13/cast"
	"strconv"
)

var (
	// ToInt
	/*
	   e.g.
	   (nil) => 0
	*/
	ToInt = cast.ToInt

	// ToIntE
	/*
	   e.g.
	   (nil) 	=> 0 <nil>
	   (false)	=> 0 <nil>
	   (true)	=> 1 <nil>
	   ("")		=> 0 unable to cast "" of type string to int64
	*/
	ToIntE = cast.ToIntE

	ToInt8 = cast.ToInt8

	ToInt8E = cast.ToInt8E

	ToInt16 = cast.ToInt16

	ToInt16E = cast.ToInt16E

	ToInt32 = cast.ToInt32

	ToInt32E = cast.ToInt32E

	ToInt64 = cast.ToInt64

	ToInt64E = cast.ToInt64E

	// StringToInt 类型转换: string => int
	StringToInt = strconv.Atoi
)

// StringToIntWithDefault 类型转换: string => int
func StringToIntWithDefault(str string, def int) int {
	i, err := StringToInt(str)
	if err != nil {
		return def
	}
	return i
}
