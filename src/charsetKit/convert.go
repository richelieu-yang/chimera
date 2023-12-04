package charsetKit

import "github.com/gogf/gf/v2/encoding/gcharset"

var (
	// Convert 转换字符串的编码（字符集的编码）
	/*
	   支持的字符集（charset）: UTF-8、GBK、Big5、ISO-* 等，更多详见: https://goframe.org/pages/viewpage.action?pageId=1114178.
	*/
	Convert func(dstCharset string, srcCharset string, src string) (dst string, err error) = gcharset.Convert

	ToUTF8 func(srcCharset string, src string) (dst string, err error) = gcharset.ToUTF8

	UTF8To func(dstCharset string, src string) (dst string, err error) = gcharset.UTF8To
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
