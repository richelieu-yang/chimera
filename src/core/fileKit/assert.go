package fileKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

// AssertExist
/*
@param path 文件（或目录）的路径
*/
func AssertExist(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exist(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertNotExistOrIsFile
/*
通过的情况: 	不存在 || 存在但是个文件
不通过的情况:	存在但是个目录
*/
func AssertNotExistOrIsFile(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if Exist(path) && IsDir(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertNotExistOrIsDir
/*
通过的情况: 	不存在 || 存在但是个目录
不通过的情况:	存在但是个文件
*/
func AssertNotExistOrIsDir(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if Exist(path) && IsFile(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func AssertExistAndIsFile(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exist(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsDir(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func AssertExistAndIsDir(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exist(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsFile(path) {
		return errorKit.NewSkipf(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}
