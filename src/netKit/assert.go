package netKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertValidPort(port int) error {
	if !IsValidPort(port) {
		return errorKit.NewSkip(1, "[%s] port(%d) is invalid", funcKit.GetFuncName(1), port)
	}
	return nil
}
