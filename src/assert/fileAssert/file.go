package fileAssert

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/file/fileKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/funcKit"
)

// AssertExist
/*
@param path 文件（或目录）的路径
*/
func AssertExist(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !fileKit.Exist(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
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
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if fileKit.Exist(path) && fileKit.IsDir(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) exists but is a directory", funcKit.GetFuncName(1), path)
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
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if fileKit.Exist(path) && fileKit.IsFile(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) exists but is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func AssertExistAndIsFile(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !fileKit.Exist(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if fileKit.IsDir(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) exists but is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func AssertExistAndIsDir(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !fileKit.Exist(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if fileKit.IsFile(path) {
		return errorKit.SimpleWithExtraSkip(1, "[%s] path(%s) exists but is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}
