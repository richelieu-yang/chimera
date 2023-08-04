package pathKit

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Walk 遍历目录.
/*
PS:
(1) 包含传参root
(2) filepath.Walk 和 filepath.WalkDir 的区别:
*/
var Walk func(root string, fn filepath.WalkFunc) error = filepath.Walk

var WalkDir func(root string, fn fs.WalkDirFunc) error = filepath.WalkDir

// ReadDir 遍历目录.
/*
@param name 目录路径

PS:
(1) 获取 传参name 目录下的所有文件（或目录）；
(2) 不包含子目录下的文件（或目录）.
*/
var ReadDir func(name string) ([]os.DirEntry, error) = os.ReadDir
