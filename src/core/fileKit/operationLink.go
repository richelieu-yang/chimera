package fileKit

import "os"

// CreateSoftLink 创建软链接
/*
Golang 之 文件硬连接 与 软连接: https://blog.csdn.net/icebergliu1234/article/details/109208030

@param src	源文件
@param dest	生成链接的位置
*/
func CreateSoftLink(src, dest string) error {
	return os.Symlink(src, dest)
}

// CreateHardLink 创建软链接
/*
Golang 之 文件硬连接 与 软连接: https://blog.csdn.net/icebergliu1234/article/details/109208030

@param src	源文件
@param dest	生成链接的位置
*/
func CreateHardLink(src, dest string) error {
	return os.Link(src, dest)
}
