package strKit

import (
	"github.com/spf13/cast"
	"strconv"
)

func IntToString(i int) string {
	//return ToString(i)
	return strconv.Itoa(i)
}

// ToString
/*
@param obj 支持的类型：time.Duration...

e.g.
(nil) => ""
*/
func ToString(obj interface{}) string {
	return ToStringWithDefault(obj, "")
}

func ToStringWithDefault(obj interface{}, def string) string {
	str, _ := ToStringWithDefaultE(obj, def)
	return str
}

func ToStringWithDefaultE(obj interface{}, def string) (string, error) {
	if obj == nil {
		return def, nil
	}
	rst, err := cast.ToStringE(obj)
	if err != nil {
		return def, err
	}
	return rst, nil
}
