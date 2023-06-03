package charsetKit

import "github.com/gogf/gf/v2/encoding/gcharset"

// Convert 转换字符串的编码（字符集的编码）
/*
支持的字符集（charset）: "UTF-8"、"GBK"、"Big5"等，更多详见: https://goframe.org/pages/viewpage.action?pageId=1114178.
*/
func Convert(dstCharset string, srcCharset string, src string) (dst string, err error) {
	return gcharset.Convert(dstCharset, src, src)
}

// IsSupported 是否支持 指定字符集 ？
func IsSupported(charset string) bool {
	return gcharset.Supported(charset)
}

func ToUTF8(srcCharset string, src string) (dst string, err error) {
	return gcharset.ToUTF8(srcCharset, src)
}

func UTF8To(dstCharset string, src string) (dst string, err error) {
	return gcharset.UTF8To(dstCharset, src)
}
