package fileKit

import (
	"syscall"
)

// IsHidden 文件（或目录）是否隐藏？
/*
如何在 Go 中检测文件夹中的隐藏文件 - 跨平台方法
  https://www.likecs.com/ask-919454.html#sc=1368.5

PS:
(1) 传参path 对应的文件或目录必须存在，否则返回error.

@param path 文件或目录的路径（绝对||相对）
*/
func IsHidden(path string) (bool, error) {
	if err := AssertExist(path); err != nil {
		return false, err
	}

	pointer, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return false, err
	}
	attributes, err := syscall.GetFileAttributes(pointer)
	if err != nil {
		return false, err
	}
	return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
}
