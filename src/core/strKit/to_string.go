package strKit

import (
	"github.com/spf13/cast"
	"strconv"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

// ToString
/*
@param obj 支持的类型：time.Duration...

e.g.
(nil) => ""
*/
func ToString(obj interface{}) string {
	return cast.ToString(obj)
}

func ToStringE(obj interface{}) (string, error) {
	return cast.ToStringE(obj)
}

//func ParseInterfaceToString(obj interface{}, def string) (string, error) {
//	if obj == nil {
//		return def, nil
//	}
//
//	cast.ToString()
//
//	rst, err := cast.ToStringE(obj)
//	if err != nil {
//		return def, err
//	}
//	return rst, nil
//}
