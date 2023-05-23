// Package randomKit
/*
有两个标准库("math/rand"、"crypto/rand")，建议使用第一个，更加全面.
*/
package randomKit

// Int
/*
参考: GoFrame中的 grand.N().

@return 范围: [min, max)
*/
func Int(min, max int) int {
	if min >= max {
		return min
	}
	return r.Intn(max-min) + min
}
