package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	path := ""

	fmt.Println(IsHidden(path))
}

// IsHidden 文件（或目录）是否隐藏？
/*
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
