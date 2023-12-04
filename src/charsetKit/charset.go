package charsetKit

import (
	"github.com/gogf/gf/v2/encoding/gcharset"
)

// IsSupported 是否支持 指定字符集 ？
var IsSupported func(charset string) bool = gcharset.Supported
