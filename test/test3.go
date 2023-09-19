package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"syscall"
)

func main() {
	path := "/Users/richelieu/Downloads/a.txt"
	fmt.Println("ReadWrite:", HaveReadWriteAccess(path))
	fmt.Println("Read:", HaveReadAccess(path))
	fmt.Println("Write:", HaveWriteAccess(path))
}

// HaveReadWriteAccess 当前用户 对指定文件是否有 读写权限 ?
/*
	PS:
	(1) 返回值为nil: 文件不存在 || 无指定权限;
	(2) 只要缺一种权限，返回值就不为nil.
*/
func HaveReadWriteAccess(path string) error {
	if err := fileKit.AssertExist(path); err != nil {
		return err
	}
	return syscall.Access(path, syscall.O_RDWR)
}

func HaveWriteAccess(path string) error {
	if err := fileKit.AssertExist(path); err != nil {
		return err
	}
	return syscall.Access(path, syscall.O_WRONLY)
}

func HaveReadAccess(path string) error {
	if err := fileKit.AssertExist(path); err != nil {
		return err
	}
	return syscall.Access(path, syscall.O_RDONLY)
}
