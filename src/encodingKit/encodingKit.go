// Package encodingKit
/*
UTF-8 是golang的原生编码方式.

参考：
Go-文本编码的转换处理 使用golang.org/x/text库
	https://blog.csdn.net/gaoluhua/article/details/109128154
go语言中的UTF-8与GBK编码转换
	https://wenku.baidu.com/view/d6a6e1c2561810a6f524ccbff121dd36a32dc480.html
*/
package encodingKit

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

// Utf8ToGbkString UTF-8 => GBK
func Utf8ToGbkString(str string) (string, int, error) {
	return transform.String(simplifiedchinese.GBK.NewEncoder(), str)
}

// GbkToUtf8String GBK => UTF-8
func GbkToUtf8String(str string) (string, int, error) {
	return transform.String(simplifiedchinese.GBK.NewDecoder(), str)
}

// Utf8ToGbk UTF-8 => GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// GbkToUtf8 GBK => UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	//第二个参数为“transform.Transformer”接口，simplifiedchinese.GBK.NewDecoder()包含了该接口
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
