package fileKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"os"
)

// Open 此方法对 os.Open 进行了封装
/*
@param path 文件（或目录）的路径

PS:
(1) 对于os.Open()，如果传参对应的文件不存在，将返回error.
(2) os.Open() 是以"只读"权限打开.
*/
func Open(path string) (*os.File, error) {
	if !Exist(path) {
		return nil, errorKit.Simple("path(%s) does not exist", path)
	}

	return os.Open(path)
}
