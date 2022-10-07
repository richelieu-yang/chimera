package assertKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
)

func NotNil(obj interface{}, variableName string) error {
	if obj == nil {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s is nil", variableName)
	}
	return nil
}

func IsNil(obj interface{}, variableName string) error {
	if obj != nil {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] variable %s isn't nil", variableName)
	}
	return nil
}
