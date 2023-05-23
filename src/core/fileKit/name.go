package fileKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"os"
	"path"
	"path/filepath"
)

// GetName 获取文件（或目录）的名称（带后缀）.
/*
PS:
(1) 文件可以不存在；
(2) 不要使用 path.Base()，不然在Windows环境下有问题.

@param pathOfFile 文件名 或 路径（相对||绝对）

e.g.
("") 		=> "."
(" ") 		=> " "（1个空格）
("  ") 		=> "  "（2个空格）
(".") 		=> "."
("./") 		=> "."
("../") 	=> ".."

e.g.1
("test.log") 						=> "test.log"
(""/Users/richelieu/Downloads"") 	=> "Downloads"
("c:/a/c/aaa.doc") 					=> "aaa.doc"

e.g.2
("/Users/richelieu/START")	=> "START"
("/Users/richelieu/START/")	=> "START"
*/
func GetName(nameOrPath string) string {
	//return path.Base(nameOrPath)
	return filepath.Base(nameOrPath)
}

// GetPrefix	文件名前缀
/**
PS: 文件可以不存在.
参考：https://zhuanlan.zhihu.com/p/80403583

@param pathOfFile 文件名 或 文件路径（相对||绝对）

e.g.
"d:/t/test.log" 	=> "test"
"d:/t/test" 		=> "test"
*/
func GetPrefix(nameOrPath string) string {
	name := GetName(nameOrPath)
	suffix := GetSuffix(name)
	// 截取字符串，0 <= 下标 < len(name) - len(suffix)
	return name[:len(name)-len(suffix)]
}

// GetSuffix 返回文件名后缀（带"."; 小写）
/*
PS: 文件可以不存在.

@param pathOfFile 文件名 或 文件路径（相对||绝对）

e.g.
("d:/t/test.LOG")	=> ".log"
("d:/t/test") 		=> ""
*/
func GetSuffix(nameOrPath string) string {
	return strKit.ToLower(path.Ext(nameOrPath))
}

// GetSuffixWithoutDot
/*
@return 文件名后缀（小写；不带"."）
*/
func GetSuffixWithoutDot(nameOrPath string) string {
	str := strKit.ToLower(path.Ext(nameOrPath))
	if strKit.StartWith(str, ".") {
		return strKit.SubAfter(str, 1)
	}
	return str
}

// Rename 重命名文件（或目录）
/*
PS:
(1) 重命名的同时，也能 将该文件移动到别的目录下；
(2) 重命名目录，该目录下 有没有文件或目录 不会有影响，正常能成功；
(3) 重命名文件，如果 newPath 对应的是一个已经存在的文件，将覆盖那个文件（并不是加到最后面）；
(4) 不管 oldPath，如果 newPath 对应的是一个已经存在的目录，将返回error(e.g. rename /Users/richelieu/Downloads/1 /Users/richelieu/Downloads/2: file exists).
*/
func Rename(oldPath, newPath string) error {
	if err := AssertExist(oldPath); err != nil {
		return err
	}

	// 创建newFilePath父路径的目录，以防 os.Rename() 返回error（e.g.rename /Users/richelieu/Downloads/a.txt /Users/richelieu/Downloads/1/2/3/b.txt: no such file or directory）
	if err := MkParentDirs(newPath); err != nil {
		return err
	}
	return os.Rename(oldPath, newPath)
}
