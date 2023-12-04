package charsetKit

import (
	"github.com/gogf/gf/v2/encoding/gcharset"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

// Convert 转换字符串的编码（字符集的编码）
/*
支持的字符集（charset）: UTF-8、GBK、Big5、ISO-* 等，更多详见: https://goframe.org/pages/viewpage.action?pageId=1114178.
*/
var Convert func(dstCharset string, srcCharset string, src string) (dst string, err error) = gcharset.Convert

// IsSupported 是否支持 指定字符集 ？
var IsSupported func(charset string) bool = gcharset.Supported

var ToUTF8 func(srcCharset string, src string) (dst string, err error) = gcharset.ToUTF8

var UTF8To func(dstCharset string, src string) (dst string, err error) = gcharset.UTF8To

var DetermineEncoding func(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) = charset.DetermineEncoding
