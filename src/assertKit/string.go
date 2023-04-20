package assertKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

func NotEmpty(str, variableName string) error {
	if strKit.IsEmpty(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable(name: %s) is empty", variableName)
	}
	return nil
}

func NotBlank(str, variableName string) error {
	if strKit.IsBlank(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable(name: %s) is blank", variableName)
	}
	return nil
}
