package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

var (
	// Open （只读模式）打开文件/目录.
	/*
		PS:
	*/
	Open func(path string) (*os.File, error) = gfile.Open

	// Create
	/*
		PS:
	*/
	Create func(path string) (*os.File, error) = gfile.Create

	// OpenFile
	/*
		@param flag 详见".info"
		@param perm
	*/
	OpenFile func(path string, flag int, perm os.FileMode) (*os.File, error) = gfile.OpenFile
)

//// Open 以"只读权限"打开文件（或目录）.
///*
//@param path 文件（或目录）的路径
//
//PS:
//(1) 对于os.Open()，如果传参对应的文件不存在，将返回error.
//(2) os.Open() 是以"只读"权限打开.
//*/
//func Open(path string) (*os.File, error) {
//	if err := AssertExist(path); err != nil {
//		return nil, err
//	}
//
//	return os.Open(path)
//}
