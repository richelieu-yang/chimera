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

func ReadFileString(filePath string) (string, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return "", err
	}

	data, err := os.ReadFile(filePath)
	return string(data), err
}

// ReadLuaFile 按行读取 .lua文件 的内容.
/*
@param path .lua文件的路径
*/
func ReadLuaFile(path string) (string, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return "", err
	}

	var buffer = bytes.Buffer{}
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		line = strKit.Trim(line)
		// 忽略"空行"和"注释行"
		if strKit.IsEmpty(line) || strKit.StartWith(line, "--") {
			continue
		}
		// 加个空格
		buffer.WriteString(line + " ")

		//line := scan.Bytes()
		//buffer.Write(line)
	}
	if err := scan.Err(); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
