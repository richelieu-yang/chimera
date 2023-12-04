package charsetKit

import (
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

// DetermineEncoding 编码推断.
/*
PS: 检测GBK文本有问题，推断出来的是"windows-1252".
*/
var DetermineEncoding func(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) = charset.DetermineEncoding

// Detect 检测文本的编码.
func Detect(data []byte) (charset string, err error) {
	var r *chardet.Result
	r, err = chardet.NewTextDetector().DetectBest(data)
	if err != nil {
		return
	}
	return r.Charset, nil
}
