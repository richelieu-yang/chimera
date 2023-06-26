package interfaceKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertNotNil(obj interface{}, name string) error {
	if obj == nil {
		return errorKit.NewSkip(1, "[%s] param(name: %s) == nil", funcKit.GetFuncName(1), name)
	}
	return nil
}
