package fileKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"os"
	"path/filepath"
)

// MkDirs 为目录路径，创建（一级或多级）目录.
/*
PS:
(1) 如果目录已经存在，将返回nil；
(2) 如果 传参dirPath 对应的是个已存在的文件，将返回error（"mkdir {xxx}: not a directory"）.

@param dirPath 相对路径 || 绝对路径

e.g.
("i:/test/test.exe") 	=> 	路径没问题且目录不存在的情况下，会在i盘创建"test"、"test.exe"两个目录
("i:/test1/test2/")		=>	路径没问题且目录不存在的情况下，会在i盘创建"test1"、"test2"两个目录
("")					=>	nil（什么都不会做）
(".")					=>	nil（什么都不会做）
*/
func MkDirs(dirPath string) error {
	if strKit.IsEmpty(dirPath) {
		return nil
	}
	return os.MkdirAll(dirPath, os.ModePerm)
}

// MkParentDirs 为文件的父路径，创建（一级或多级）目录.
/*
e.g.
("")	=> nil
(".")	=> nil
*/
func MkParentDirs(filePath string) error {
	// Richelieu: 为防止 import cycle
	//dirPath := pathKit.GetParentDir(filePath)

	dirPath := filepath.Dir(filePath)
	return MkDirs(dirPath)
}
