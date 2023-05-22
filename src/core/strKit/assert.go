package strKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

func AssertNotEmpty(str string) error {
	if IsEmpty(str) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] str is empty", funcKit.GetFuncName(1))
	}
	return nil
}

func AssertNotBlank(str string) error {
	if IsBlank(str) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] str(value: %s, length: %d) is blank", funcKit.GetFuncName(1), str, len(str))
	}
	return nil
}
