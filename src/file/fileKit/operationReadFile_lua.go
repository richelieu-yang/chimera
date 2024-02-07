package fileKit

import (
	"bufio"
	"bytes"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

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
