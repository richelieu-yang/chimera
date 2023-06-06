package mapKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertNotEmpty[K comparable, V any](m map[K]V, name string) error {
	if len(m) == 0 {
		if m == nil {
			return errorKit.NewSkipf(1, "[%s] param(name: %s, type: %s) == nil",
				funcKit.GetFuncName(1), name, "map")
		}
		return errorKit.NewSkipf(1, "[%s] param(name: %s, type: %s) is empty",
			funcKit.GetFuncName(1), name, "map")
	}
	return nil
}
