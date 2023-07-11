package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"os"
)

var (
	// Open （只读模式）打开文件/目录.
	/*
		PS: flag == O_RDONLY.
	*/
	Open func(path string) (*os.File, error) = gfile.Open

	// OpenFile （以指定 flag 和 perm）打开文件/目录.
	/*
		@param flag 详见".info"
		@param perm	(1) 可以参考 "fileKit/consts.go"
					(2) e.g.0666 || os.ModePerm ...
	*/
	OpenFile func(path string, flag int, perm os.FileMode) (*os.File, error) = gfile.OpenFile
)

// Create 创建文件（目录不行）.
/*
	PS:
	(1) flag == O_RDWR|O_CREATE|O_TRUNC
	(2) 读写权限；
	(3) path不存在，会创建；
	(4) path存在 && 是文件，会清空该文件内容；
	(5) path存在 && 是目录，会返回error.
*/
func Create(path string) (*os.File, error) {
	//return gfile.Create(path)

	if err := MkParentDirs(path); err != nil {
		return nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		err = errorKit.Wrap(err, `os.Create failed for name "%s"`, path)
	}
	return file, err
}

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
