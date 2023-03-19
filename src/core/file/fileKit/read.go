package fileKit

import (
	"github.com/richelieu42/chimera/src/crypto/base64Kit"
	"os"
)

// ReadFile 读取文件的数据.
/*
PS:
(1) ioutil.ReadFile() 比 ioutil.ReadAll() 性能好，特别是大文件；
(2) 编码必须为"UTF-8"！！！

@param path 文件的路径（不能是目录的路径）
*/
func ReadFile(filePath string) ([]byte, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}
	return os.ReadFile(filePath)
}

func ReadFileToString(filePath string) (string, error) {
	data, err := ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadFileToBase64
/*
!!!: 如果想实现: 图片 => base64字符串，请使用 imageKit.ConvertImageToBase64().
*/
func ReadFileToBase64(filePath string) (string, error) {
	data, err := ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return base64Kit.EncodeToString(data), nil
}
