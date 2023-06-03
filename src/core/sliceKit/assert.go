package sliceKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

// AssertNotEmpty
/*
合法的传参:
(1) s != nil
(2) len(s) > 0
*/
func AssertNotEmpty[T any](s []T) error {
	if len(s) == 0 {
		if s == nil {
			return errorKit.NewSkipf(1, "[%s] s == nil", funcKit.GetFuncName(1))
		}
		return errorKit.NewSkipf(1, "[%s] len(s) == 0", funcKit.GetFuncName(1))
	}
	return nil
}
