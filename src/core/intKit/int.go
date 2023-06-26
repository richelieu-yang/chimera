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
	ToInt func(i interface{}) int = cast.ToInt

	// ToIntE
	/*
	   e.g.
	   (nil) 	=> 0 <nil>
	   (false)	=> 0 <nil>
	   (true)	=> 1 <nil>
	   ("")		=> 0 unable to cast "" of type string to int64
	*/
	ToIntE func(i interface{}) (int, error) = cast.ToIntE

	ToInt8 func(i interface{}) int8 = cast.ToInt8

	// ToInt8E
	/*
		e.g. 将string转换为整型，建议还是用 StringToInt
			fmt.Println(cast.ToInt8E("07")) // 7 <nil>
			fmt.Println(cast.ToInt8E("08")) // 0 unable to cast "08" of type string to int64

			fmt.Println(strconv.Atoi("07")) // 7 <nil>
			fmt.Println(strconv.Atoi("08")) // 8 <nil>
	*/
	ToInt8E func(i interface{}) (int8, error) = cast.ToInt8E

	ToInt16 func(i interface{}) int16 = cast.ToInt16

	ToInt16E func(i interface{}) (int16, error) = cast.ToInt16E

	ToInt32 func(i interface{}) int32 = cast.ToInt32

	ToInt32E func(i interface{}) (int32, error) = cast.ToInt32E

	ToInt64 func(i interface{}) int64 = cast.ToInt64

	ToInt64E func(i interface{}) (int64, error) = cast.ToInt64E

	// StringToInt 类型转换: string => int
	StringToInt func(s string) (int, error) = strconv.Atoi
)

// StringToIntWithDefault 类型转换: string => int
func StringToIntWithDefault(str string, def int) int {
	i, err := StringToInt(str)
	if err != nil {
		return def
	}
	return i
}
