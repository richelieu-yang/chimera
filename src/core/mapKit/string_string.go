package mapKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// JoinSS map[string]string => string
func JoinSS(m map[string]string, separator, keyValueSeparator string, keyCallback func(str string) string, valueCallback func(str string) string) (rst string) {
	for key, value := range m {
		if keyCallback != nil {
			key = keyCallback(key)
		}
		if valueCallback != nil {
			value = valueCallback(value)
		}

		var tmp string
		if strKit.IsEmpty(value) {
			tmp = key
		} else {
			tmp = key + keyValueSeparator + value
		}

		if strKit.IsEmpty(rst) {
			rst += tmp
		} else {
			rst += separator + tmp
		}
	}
	return
}
