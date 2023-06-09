package interfaceKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertNotNil(obj interface{}) error {
	if obj == nil {
		return errorKit.NewSkip(1, "[%s] obj == nil", funcKit.GetFuncName(1))
	}
	return nil
}
