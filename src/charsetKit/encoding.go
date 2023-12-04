package charsetKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"unicode/utf8"
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

var IsUTF8 func(data []byte) bool = utf8.Valid

var IsUTF8String func(s string) bool = utf8.ValidString

// IsGBK
/*
PS: 在Go语言中，你可以通过检查字节序列是否落在 GBK 编码范围内来判断文本编码是否为 GBK。
*/
func IsGBK(data []byte) bool {
	if IsUTF8(data) {
		return false
	}

	i := 0
	for i < len(data) {
		if data[i] <= 0x7f {
			// 编码0~127,只有一个字节的编码，兼容ASCII码
			i++
			continue
		}
		if data[i] >= 0x81 && data[i] <= 0xfe && data[i+1] >= 0x40 && data[i+1] <= 0xfe && data[i+1] != 0xf7 {
			// 大于127的使用双字节编码，落在gbk编码范围内的字符
			i += 2
			continue
		}
		return false
	}
	return true
}
