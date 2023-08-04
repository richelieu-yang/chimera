package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
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

	// Create 创建文件（目录不行）.
	/*
		PS:
		(1) flag == O_RDWR|O_CREATE|O_TRUNC
		(2) 读写权限；
		(3) path不存在，会创建；
		(4) path存在 && 是文件，会清空该文件内容；
		(5) path存在 && 是目录，会返回error.
	*/
	Create func(path string) (*os.File, error) = gfile.Create
)
