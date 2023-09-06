package fileKit

import "os"

// WriteToFile 将数据（字节流）写到文件中.
/*
@param filePath 目标文件的路径
				(1) 不存在的话，会创建一个新的文件;
				(2) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）.
*/
func WriteToFile(data []byte, filePath string) error {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := MkParentDirs(filePath); err != nil {
		return err
	}

	return os.WriteFile(filePath, data, os.ModePerm)
}

// WriteStringToFile 将数据（字符串）写到文件中.
/*
@param filePath 目标文件的路径（不存在的话，会创建一个新的文件；存在且是个文件的话，会覆盖掉旧的（并不会加到该文件的最后面））
*/
func WriteStringToFile(str, filePath string) error {
	return WriteToFile([]byte(str), filePath)
}
