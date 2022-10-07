package fileKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

// AssertExist
/*
@param path 文件（或目录）的路径
*/
func AssertExist(path string) error {
	return assertExist(path, 1)
}

// AssertNotExistOrIsFile
/*
@return 返回nil（通过断言）的情况: 不存在 || 存在但是个文件
*/
func AssertNotExistOrIsFile(path string) error {
	if Exist(path) && IsDir(path) {
		// 此处的1是为了跳过当前函数的调用
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] path(%s) exists but is a directory", path)
	}
	return nil
}

// AssertNotExistOrIsDir
/*
@return 返回nil（通过断言）的情况: 不存在 || 存在但是个目录
*/
func AssertNotExistOrIsDir(path string) error {
	if Exist(path) && IsFile(path) {
		// 此处的1是为了跳过当前函数的调用
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] path(%s) exists but is a file", path)
	}
	return nil
}

// assertExist
/*
@param extraSkip 必须>=1
*/
func assertExist(path string, extraSkip int) error {
	if !Exist(path) {
		// 此处的1是为了跳过当前函数的调用
		return errorKit.SimpleWithExtraSkip(1+extraSkip, "[Assertion failed] path(%s) doesn't exist", path)
	}
	return nil
}

// AssertExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func AssertExistAndIsFile(path string) error {
	if err := assertExist(path, 1); err != nil {
		return err
	}
	if !IsFile(path) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] path(%s) exists but not a file", path)
	}
	return nil
}

// AssertExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func AssertExistAndIsDir(path string) error {
	if err := assertExist(path, 1); err != nil {
		return err
	}
	if !IsDir(path) {
		return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] path(%s) exists but not a directory", path)
	}
	return nil
}
