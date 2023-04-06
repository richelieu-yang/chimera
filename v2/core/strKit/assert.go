package strKit

import "github.com/richelieu42/chimera/v2/core/errorKit"

func AssertNotEmpty(str, variableName string) error {
	if IsEmpty(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s(string type) is empty", variableName)
	}
	return nil
}

func AssertNotBlank(str, variableName string) error {
	if IsBlank(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s(string type) is blank", variableName)
	}
	return nil
}
