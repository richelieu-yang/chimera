package pathKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"path/filepath"
)

// outputDir （默认的）输出目录
var outputDir string

var (
	// SelfDir returns absolute directory path of current running process(binary).
	SelfDir = gfile.SelfDir

	// MainPkgPath returns absolute file path of package main, which contains the entrance function main.
	MainPkgPath = gfile.MainPkgPath
)

// GetParentDir 返回文件（或目录）路径的父路径.
/*
PS: 类似于Java中的 getParentFile().
*/
func GetParentDir(path string) string {
	//return filepath.Dir(path)
	return gfile.Dir(path)
}

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

// VolumeName 返回分区名
/*
PS: 非Windows平台，将返回"".

e.g. Windows平台
("C:/a/b/c") => "C:"
*/
func VolumeName(path string) string {
	return filepath.VolumeName(path)
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

// CheckSkip 检查是否发生"路径穿越"（路径穿透）
/*
@param path	 	父路径
@param path1	子路径
@return true: 发生"路径穿越"

e.g.
("/a//b/", "/a//b//../c.docx")	=> true
("/a//b/", "//a//b///c.docx")	=> false
*/
func CheckSkip(parent, path string) bool {
	parent = Join(parent)
	path = Join(path)

	return !strKit.StartWith(path, parent)
}

// GetOutputPath 获取输出目录的路径
func GetOutputPath(timeStr string) (string, error) {
	if strKit.IsEmpty(outputDir) {
		if strKit.IsEmpty(timeStr) {
			timeStr = timeKit.FormatCurrentTime(timeKit.FormatFileName)
		}
		tmp := Join(GetWorkingDir(), "out", "BootTime_"+timeStr)
		if err := fileKit.MkDirs(tmp); err != nil {
			return "", err
		}
		outputDir = tmp
	}
	return outputDir, nil
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

// EvalSymlinks 返回链接文件的实际路径
/*
@param path e.g."1.lnk"
*/
func EvalSymlinks(path string) (string, error) {
	return filepath.EvalSymlinks(path)
}
