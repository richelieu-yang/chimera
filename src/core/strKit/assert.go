package strKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertStringNotEmpty(str string) error {
	if IsEmpty(str) {
		return errorKit.NewSkipf(1, "[%s] str is empty", funcKit.GetFuncName(1))
	}
	return nil
}

func AssertStringNotBlank(str string) error {
	if IsBlank(str) {
		return errorKit.NewSkipf(1, "[%s] str(%s) is blank", funcKit.GetFuncName(1), str)
	}
	return nil
}
