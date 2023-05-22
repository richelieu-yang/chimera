package interfaceAssert

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertNotNil(obj interface{}) error {
	if obj == nil {
		return errorKit.Simple("[%s] obj == nil", funcKit.GetFuncName(1))
	}
	return nil
}
