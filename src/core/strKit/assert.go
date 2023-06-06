package strKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

func AssertNotEmpty(str string) error {
	if IsEmpty(str) {
		return errorKit.NewSkipf(1, "[%s] str is empty", funcKit.GetFuncName(1))
	}
	return nil
}

func AssertNotBlank(str string) error {
	if IsBlank(str) {
		return errorKit.NewSkipf(1, "[%s] str(%s) is blank", funcKit.GetFuncName(1), str)
	}
	return nil
}
