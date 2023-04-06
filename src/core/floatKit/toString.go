// Package floatKit
package floatKit

import (
	"strconv"
)

// FormatFloat32ToString 类型转换: float32 => string
/*
PS: 传参可参考"Golang.docx"中的"strconv标准库".
*/
func FormatFloat32ToString(f float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, 32)
}

// FormatFloat64ToString 类型转换: float64 => string
/*
PS: 传参可参考"Golang.docx"中的"strconv标准库".
*/
func FormatFloat64ToString(f float64, fmt byte, prec int) string {
	return strconv.FormatFloat(f, fmt, prec, 64)
}
