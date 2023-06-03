package urlKit

import "github.com/richelieu-yang/chimera/v2/src/core/strKit"

func ToBodyString(postParams map[string]string) string {
	var str string

	for k, v := range postParams {
		if strKit.IsNotEmpty(str) {
			str += "&"
		}
		str += k + "=" + EncodeURIComponent(v)
	}
	return str
}
