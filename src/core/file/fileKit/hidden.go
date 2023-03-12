//go:build !windows

package fileKit

import (
	"path/filepath"
	"strings"
)

// IsHidden 文件（或目录）是否隐藏？
/*
如何在 Go 中检测文件夹中的隐藏文件 - 跨平台方法
	https://www.likecs.com/ask-919454.html#sc=1368.5

流程:
(1) 获取文件名（以防传参为路径）
(2) 判断文件名是否以"."开头

@param 文件（或目录）的 path 或 name

e.g.
("") => false, nil
*/
func IsHidden(path string) (bool, error) {
	name := filepath.Base(path)

	switch name {
	case ".":
		fallthrough
	case "..":
		return false, nil
	default:
		return strings.HasPrefix(name, "."), nil
	}
}
