package charsetKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
)

//import (
//	"bytes"
//	"golang.org/x/text/encoding/simplifiedchinese"
//	"golang.org/x/text/transform"
//	"io"
//)
//
//// Utf8ToGbkString UTF-8 => GBK
//func Utf8ToGbkString(str string) (string, int, error) {
//	return transform.ToDsnString(simplifiedchinese.GBK.NewEncoder(), str)
//}
//
//// GbkToUtf8String GBK => UTF-8
//func GbkToUtf8String(str string) (string, int, error) {
//	return transform.ToDsnString(simplifiedchinese.GBK.NewDecoder(), str)
//}
//
//// Utf8ToGbk UTF-8 => GBK
//func Utf8ToGbk(s []byte) ([]byte, error) {
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
//	d, e := io.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}
//
//// GbkToUtf8 GBK => UTF-8
//func GbkToUtf8(s []byte) ([]byte, error) {
//	//第二个参数为“transform.Transformer”接口，simplifiedchinese.GBK.NewDecoder()包含了该接口
//	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
//	d, e := io.ReadAll(reader)
//	if e != nil {
//		return nil, e
//	}
//	return d, nil
//}

// DetermineEncoding 编码推断.
/*
PS: 检测GBK文本有问题，推断出来的是 "windows-1252".
*/
var DetermineEncoding func(content []byte, contentType string) (e encoding.Encoding, name string, certain bool) = charset.DetermineEncoding

// Detect 检测文本的编码.
/*
PS: 检测GBK文本有问题，推断出来的是 "ISO-8859-1".
*/
func Detect(data []byte) (charset string, err error) {
	if err = interfaceKit.AssertNotNil(data, "data"); err != nil {
		return
	}

	var r *chardet.Result
	r, err = chardet.NewTextDetector().DetectBest(data)
	if err != nil {
		return
	}
	return r.Charset, nil
}
