package boolKit

import "strconv"

// FormatBoolToString 类型转换: bool => string
func FormatBoolToString(b bool) string {
	return strconv.FormatBool(b)
}
