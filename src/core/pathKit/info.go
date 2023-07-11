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
	SelfDir func() string = gfile.SelfDir

	// MainPkgPath returns absolute file path of package main, which contains the entrance function main.
	MainPkgPath func() string = gfile.MainPkgPath

	// ParentDir 返回文件（或目录）路径的父路径.
	/*
	   PS: 类似于Java中的 getParentFile().
	*/
	ParentDir func(path string) string = gfile.Dir

	// IsAbs 判断路径是不是绝对路径
	/*
	   PS: 传参path 对应的文件（或目录）可以不存在.

	   e.g.
	   ("./a/b/c") 	=> false
	   ("C:/a/b/c") 	=> true
	   ("/root") 		=> true
	   ("root") 		=> false
	*/
	IsAbs func(path string) bool = filepath.IsAbs

	// Abs 获取的绝对路径(传参path相对于当前路径(os.Getwd())).
	Abs func(path string) (string, error) = filepath.Abs

	// VolumeName 返回分区名
	/*
	   PS: 非Windows平台，将返回"".

	   e.g. Windows平台
	   ("C:/a/b/c") => "C:"
	*/
	VolumeName func(path string) string = filepath.VolumeName

	// GetRelativePath 获取（targetPath 相对于 basePath 的）相对路径.
	/*
	   e.g. Windows
	   ("C:/a/b", "C:/a/b/c/d/../e") => "c\e", <nil>

	   e.g.1 Mac
	   ("/usr/local", "/usr/local/go/bin")			=> "go/bin", nil
	   ("//usr////local", "/usr/local/go/bin")		=> "go/bin", nil
	   ("//usr////local", "/usr/local/go/bin/../")	=> "go", nil
	*/
	GetRelativePath func(basepath, targpath string) (string, error) = filepath.Rel

	// Clean 返回等价的最短路径（清理路径）
	/*
	   PS: 可用于去掉路径中的 "./" 、 "../" ...

	   e.g.
	   ("./1.txt") 		=> "1.txt"
	   ("/root/.././c") 	=> "/c"
	*/
	Clean func(path string) string = filepath.Clean

	// Match （正则相关）匹配文件名，完全匹配则返回true
	/*
	   e.g.
	   ("*", "a") => (true, <nil>)
	   ("*", "C:/a/b/c") => (true, <nil>)
	   ("\\b", "b") => (false, <nil>)
	*/
	Match = filepath.Match

	// Glob （正则相关）返回所有匹配的文件.
	/*
	   e.g.
	   ("d:/test/*.txt") => ([d:\test\a.txt d:\test\b.txt], <nil>)
	*/
	Glob = filepath.Glob

	// EvalSymlinks 返回链接文件的实际路径.
	EvalSymlinks = filepath.EvalSymlinks
)

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

// GetOutputPath 获取输出目录的路径.
/*
Deprecated: 没啥用（因为每个项目都有其特性）.
*/
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
