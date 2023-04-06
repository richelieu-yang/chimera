package boolKit

import "strconv"

// ToString 类型转换: bool => string
/*
e.g.
(true) => 	"true"
(false) => 	"false"
*/
func ToString(b bool) string {
	return strconv.FormatBool(b)
}
