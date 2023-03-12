package fileKit

import "syscall"

// IsHidden
/*
如何在 Go 中检测文件夹中的隐藏文件 - 跨平台方法
	https://www.likecs.com/ask-919454.html#sc=1368.5
*/
func IsHidden(filename string) (bool, error) {
	pointer, err := syscall.UTF16PtrFromString(filename)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
