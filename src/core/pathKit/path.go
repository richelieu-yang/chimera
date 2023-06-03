// Package pathKit
/*
@Author Richelieu
@Description  主要是对"path"、"path/filepath"的封装.

PS:
filepath标准库的使用可以参考: https://www.cnblogs.com/jkko123/p/6923962.html
*/
package pathKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"path/filepath"
)

// outputDir （默认的）输出目录
var outputDir string

// IsAbs 判断路径是不是绝对路径
/*
PS: 传参path 对应的文件（或目录）可以不存在.

e.g.
("./a/b/c") 	=> false
("C:/a/b/c") 	=> true
("/root") 		=> true
("root") 		=> false
*/
func IsAbs(path string) bool {
	return filepath.IsAbs(path)
}

// Match （正则相关）匹配文件名，完全匹配则返回true
/*
e.g.
("*", "a") => (true, <nil>)
("*", "C:/a/b/c") => (true, <nil>)
("\\b", "b") => (false, <nil>)
*/
func Match(pattern, name string) (matched bool, err error) {
	return filepath.Match(pattern, name)
}

// Glob （正则相关）返回所有匹配的文件.
/*
e.g.
("d:/test/*.txt") => ([d:\test\a.txt d:\test\b.txt], <nil>)
*/
func Glob(pattern string) (matches []string, err error) {
	return filepath.Glob(pattern)
}

// GetOutputPath 获取输出目录的路径
func GetOutputPath(timeStr string) (string, error) {
	if strKit.IsEmpty(outputDir) {
		if strKit.IsEmpty(timeStr) {
			timeStr = timeKit.FormatCurrentTime(timeKit.FormatFileName)
		}

		workingDir, err := GetWorkingDir()
		if err != nil {
			return "", err
		}
		tmp := Join(workingDir, "out", "BootTime_"+timeStr)
		if err := fileKit.MkDirs(tmp); err != nil {
			return "", err
		}
		outputDir = tmp
	}
	return outputDir, nil
}

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

// VolumeName 返回分区名
/*
PS: 非Windows平台，将返回"".

e.g. Windows平台
("C:/a/b/c") => "C:"
*/
func VolumeName(path string) string {
	return filepath.VolumeName(path)
}

// GetRelativePath 获取（targetPath 相对于 basePath 的）相对路径.
/*
e.g. Windows
("C:/a/b", "C:/a/b/c/d/../e") => "c\e", <nil>

e.g.1 Mac
("/usr/local", "/usr/local/go/bin")			=> "go/bin", nil
("//usr////local", "/usr/local/go/bin")		=> "go/bin", nil
("//usr////local", "/usr/local/go/bin/../")	=> "go", nil
*/
func GetRelativePath(basePath, targetPath string) (string, error) {
	return filepath.Rel(basePath, targetPath)
}

// GetAbsolutePath 获取的绝对路径(传参path相对于当前路径(os.Getwd())).
func GetAbsolutePath(path string) (string, error) {
	return filepath.Abs(path)
}

// IsAbsolutePath 是否绝对路径？
/*
e.g.
("") 		=> false
("./a/b/c")	=> false
("/a/b/c")	=> true
*/
func IsAbsolutePath(path string) bool {
	return filepath.IsAbs(path)
}

// Clean 返回等价的最短路径（清理路径）
/*
PS: 可用于去掉路径中的 "./" 、 "../" ...

e.g.
("./1.txt") 		=> "1.txt"
("/root/.././c") 	=> "/c"
*/
func Clean(path string) string {
	return filepath.Clean(path)
}

// EvalSymlinks 返回链接文件的实际路径
/*
@param path e.g."1.lnk"
*/
func EvalSymlinks(path string) (string, error) {
	return filepath.EvalSymlinks(path)
}

// GetParentDir 返回文件（或目录）路径的父路径.
/*
PS: 类似于Java中的 getParentFile().

@param nameOrPath 文件名或文件路径

e.g.
("")			=> "."
(".")			=> "."
("yozo.eio") 	=> "."
("/")			=> "/"

e.g.1 Mac
("./a/b/c")		=> "a/b"
("C:/a/b/c")	=> "C:/a/b"
*/
func GetParentDir(nameOrPath string) string {
	return filepath.Dir(nameOrPath)
}
