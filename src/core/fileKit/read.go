package fileKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
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
	if !Exist(filePath) {
		return nil, errorKit.Simple("filePath doesn't exist", filePath)
	}
	if !IsFile(filePath) {
		return nil, errorKit.Simple("filePath isn't a file", filePath)
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
