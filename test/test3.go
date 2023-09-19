package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"syscall"
)

// HaveReadWriteAccess 当前用户对指定文件是否有 读写权限 ?
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
