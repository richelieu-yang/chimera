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
e.g.
(nil) => ""
*/
func ToString(obj interface{}) string {
	return cast.ToString(obj)
}

// ToStringE
/*
e.g.
(nil) => "", nil
*/
func ToStringE(obj interface{}) (string, error) {
	return cast.ToStringE(obj)
}
