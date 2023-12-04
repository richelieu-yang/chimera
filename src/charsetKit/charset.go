package charsetKit

import (
	"github.com/gogf/gf/v2/encoding/gcharset"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

// IsSupported 是否支持 指定字符集 ？
var IsSupported func(charset string) bool = gcharset.Supported

var DetermineEncoding func(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) = charset.DetermineEncoding
