package base64Kit

import (
	"github.com/gogf/gf/v2/encoding/gbase64"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

func Encode(src []byte) []byte {
	return gbase64.Encode(src)
}

// EncodeFile
/*
@param path 文件路径（不支持目录路径）
*/
func EncodeFile(path string) ([]byte, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return nil, err
	}

	return gbase64.EncodeFile(path)
}

// EncodeFileToString 文件 => base64字符串
func EncodeFileToString(path string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(path); err != nil {
		return "", err
	}

	return gbase64.EncodeFileToString(path)
}

var EncodeString func(src string) string = gbase64.EncodeString

var EncodeToString func(src []byte) string = gbase64.EncodeToString
