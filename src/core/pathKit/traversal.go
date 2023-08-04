package pathKit

import (
	"os"
	"path/filepath"
)

// Walk 遍历目录.
/*
PS: 包含传参root
*/
var Walk func(root string, fn filepath.WalkFunc) error = filepath.Walk

// ReadDir 遍历目录.
/*
@param name 目录路径

PS:
(1) 获取 传参name 目录下的所有文件（或目录）；
(2) 不包含子目录下的文件（或目录）.
*/
var ReadDir func(name string) ([]os.DirEntry, error) = os.ReadDir
