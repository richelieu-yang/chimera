package fileKit

import (
	"os"
)

// Open 以"只读权限"打开文件（或目录）.
/*
@param path 文件（或目录）的路径

PS:
(1) 对于os.Open()，如果传参对应的文件不存在，将返回error.
(2) os.Open() 是以"只读"权限打开.
*/
func Open(path string) (*os.File, error) {
	if err := AssertExist(path); err != nil {
		return nil, err
	}

	return os.Open(path)
}
