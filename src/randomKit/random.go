// Package randomKit
/*
有两个标准库("math/rand"、"crypto/rand")，建议使用第一个，更加全面.
*/
package randomKit

import (
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
)

// Float64
/*
@param places 保留的小数位

@return [0.0, 1.0)
*/
func Float64(places int32) float64 {
	f := r.Float64()
	return floatKit.Floor(f, places)
}
