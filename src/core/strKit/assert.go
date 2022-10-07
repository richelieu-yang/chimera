package strKit

import "gitee.com/richelieu042/go-scales/src/core/errorKit"

func AssertNotEmpty(str, variableName string) error {
	if IsEmpty(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s is empty", variableName)
	}
	return nil
}

func AssertNotBlank(str, variableName string) error {
	if IsBlank(str) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s is blank", variableName)
	}
	return nil
}
