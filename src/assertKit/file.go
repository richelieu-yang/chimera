package assertKit

import (
	"github.com/richelieu42/chimera/v2/src/core/file/fileKit"
)

// Exist
/*
@param path 文件（或目录）的路径
*/
func Exist(path string) error {
	return fileKit.AssertExist(path)
}

// NotExistOrIsFile
/*
@return 返回nil（通过断言）的情况: 不存在 || 存在但是个文件
*/
func NotExistOrIsFile(path string) error {
	return fileKit.AssertNotExistOrIsFile(path)
}

// NotExistOrIsDir
/*
@return 返回nil（通过断言）的情况: 不存在 || 存在但是个目录
*/
func NotExistOrIsDir(path string) error {
	return fileKit.AssertNotExistOrIsDir(path)
}

// ExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func ExistAndIsFile(path string) error {
	return fileKit.AssertExistAndIsFile(path)
}

// ExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func ExistAndIsDir(path string) error {
	return fileKit.AssertExistAndIsDir(path)
}
