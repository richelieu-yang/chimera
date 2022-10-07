package md5Kit

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"io"
	"os"
)

// GetFromFile 获取文件的md5
func GetFromFile(path string) (string, error) {
	data, err := fileKit.ReadFile(path)
	if err != nil {
		return "", err
	}

	return Get(data), nil
}

// GetFromFile1
// Deprecated
/*
path为 空字符串 或者 对应文件不存在 时，会报错：open : The system cannot find the file specified.
*/
func GetFromFile1(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	m := md5.New()
	_, err = io.Copy(m, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(m.Sum(nil)), nil
}

func Get(s []byte) string {
	m := md5.New()
	m.Write(s)
	return hex.EncodeToString(m.Sum(nil))
}
