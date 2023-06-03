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

func EncodeFileToString(path string) (string, error) {
	return gbase64.EncodeFileToString(path)
}

func EncodeString(src string) string {
	return gbase64.EncodeString(src)
}

func EncodeToString(src []byte) string {
	return gbase64.EncodeToString(src)
}
