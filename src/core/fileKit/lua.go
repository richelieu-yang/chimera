package fileKit

import (
	"bufio"
	"bytes"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

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
