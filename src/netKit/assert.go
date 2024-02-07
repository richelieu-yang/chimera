package netKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/funcKit"
)

func AssertValidPort(port int) error {
	if !IsValidPort(int64(port)) {
		return errorKit.NewSkip(1, "[%s] port(%d) is invalid", funcKit.GetFuncName(1), port)
	}
	return nil
}
