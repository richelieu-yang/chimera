package fileKit

import (
	"bufio"
	"bytes"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

var (
	ReadLines func(file string, callback func(line string) error) error = gfile.ReadLines

	ReadLinesBytes func(file string, callback func(bytes []byte) error) error = gfile.ReadLinesBytes
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
	return string(data), err
}

// ReadFileByLine
/*
@param f 调用scan.Bytes() || scan.Text()
*/
func ReadFileByLine(filePath string, f func(scan *bufio.Scanner)) error {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return err
	}

	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return err
	}

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		// 自定义
		f(scan)
	}
	if err := scan.Err(); err != nil {
		return err
	}
	return nil
}

// ReadLuaFileToString 按行读取 .lua文件 的内容.
/*
@param path .lua文件的路径
*/
func ReadLuaFileToString(filePath string) (string, error) {
	var buffer = bytes.Buffer{}

	err := ReadFileByLine(filePath, func(scan *bufio.Scanner) {
		text := scan.Text()
		text = strKit.TrimSpace(text)

		// 忽略"空行"和"注释行"
		if strKit.IsEmpty(text) || strKit.StartWith(text, "--") {
			return
		}
		buffer.WriteString(text)
		// 加个空格
		buffer.WriteString(" ")
	})
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
