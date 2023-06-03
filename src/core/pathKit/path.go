// Package pathKit
/*
@Author Richelieu
@Description  主要是对"path"、"path/filepath"的封装.

PS:
filepath标准库的使用可以参考: https://www.cnblogs.com/jkko123/p/6923962.html
*/
package pathKit

import (
	"path/filepath"
)

// ToSlash 将路径分隔符使用"/"替换
func ToSlash(path string) string {
	return filepath.ToSlash(path)
}

// FromSlash 将路径中的"/"替换为路径分隔符
func FromSlash(path string) string {
	return filepath.FromSlash(path)
}

// Join 将多个字符串合并为一个路径（路径拼接）.
/*
PS:
(1) 不建议用"path.Join"，有问题，e.g. 传参："d:\\nacos\\", "cyy"
(2) 此方法也可用于优化路径（处理: 路径穿越、多个连续的路径分隔符...）.

e.g.
() 			=> ""
("") 		=> ""
("", "") 	=> ""
("", "a")	=> "a"

(".", "yozo.eio") => "yozo.eio"

(" ")		=> " "
(" ", "a")	=> " /a"

e.g.1	支持: 路径穿越（路径穿透）
("/a/b", "../c.docx") => "/a/c.docx"
*/
func Join(eles ...string) string {
	return filepath.Join(eles...)
}

// Split 分割路径中的目录与文件.
/*
e.g. Mac
("/Users/richelieu/Downloads/") => "/Users/richelieu/Downloads/", ""
("/Users/richelieu/Downloads") 	=> "/Users/richelieu/", "Downloads"
*/
func Split(path string) (dir, file string) {
	return filepath.Split(path)
}

// SplitList 使用"路径列表分隔符"将路径分开.
/*
os.PathListSeparator（linux下默认为':'，windows下为';'）

e.g.
("d:/a/b/c.docx") => [d:/a/b/c.docx]
("C:/windows;C:/windows/system") => [C:/windows C:/windows/system]
*/
func SplitList(path string) []string {
	return filepath.SplitList(path)
}
