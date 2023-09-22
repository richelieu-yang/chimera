package reqKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpClientKit"
)

// IsTimeoutError 是否是请求超时错误?
var IsTimeoutError func(err error) (flag bool) = httpClientKit.IsTimeoutError
