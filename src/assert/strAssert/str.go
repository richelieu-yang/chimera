package strAssert

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertStringNotEmpty(str string) error {
	if strKit.IsEmpty(str) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] str is empty", funcKit.GetFuncName(1))
	}
	return nil
}

func AssertStringNotBlank(str string) error {
	if strKit.IsBlank(str) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] str(value: %s, length: %d) is blank", funcKit.GetFuncName(1), str, len(str))
	}
	return nil
}
