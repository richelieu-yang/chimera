package fileKit

import (
	"os"
	"syscall"
)

// IsHidden 文件（或目录）是否隐藏？
/*
@param path 文件或目录的路径（绝对||相对）
*/
func IsHidden(path string) (bool, error) {
	if err := AssertExist(path); err != nil {
		return false, err
	}

	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	attributes := info.Sys().(*syscall.Win32FileAttributeData).FileAttributes
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil

	//pointer, err := syscall.UTF16PtrFromString(path)
	//if err != nil {
	//	return false, err
	//}
	//attributes, err := syscall.GetFileAttributes(pointer)
	//if err != nil {
	//	return false, err
	//}
	//return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
