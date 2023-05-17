// Package randomKit
/*
有两个标准库("math/rand"、"crypto/rand")，建议使用第一个，更加全面.
*/
package randomKit

import "github.com/richelieu42/chimera/v2/src/core/errorKit"

// Int
/*
@return 范围: [min, max)
*/
func Int(min, max int) (rst int, err error) {
	if min >= max {
		err = errorKit.Simple("min(%d) is greater than or equal to max(%d)", min, max)
		return
	}

	// rst范围: [0, max - min)
	rst = r.Intn(max - min)
	// rst范围: [min, max)
	rst += min
	return
}
