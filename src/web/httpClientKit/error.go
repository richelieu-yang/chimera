package httpClientKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

// IsTimeoutError 是否是请求超时错误?
func IsTimeoutError(err error) (flag bool) {
	if err != nil {
		if errorKit.Is(err, context.DeadlineExceeded) {
			flag = true
			return
		}
	}
	return
}
