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

@param dirPaths 目录路径s（相对路径 || 绝对路径）

e.g.
("i:/test/test.exe") 	=> 	路径没问题且目录不存在的情况下，会在i盘创建"test"、"test.exe"两个目录
("i:/test1/test2/")		=>	路径没问题且目录不存在的情况下，会在i盘创建"test1"、"test2"两个目录
("")					=>	nil（什么都不会做）
(".")					=>	nil（什么都不会做）
*/
func MkDirs(dirPaths ...string) (err error) {
	for _, dirPath := range dirPaths {
		if strKit.IsNotEmpty(dirPath) {
			err = os.MkdirAll(dirPath, os.ModePerm)
			if err != nil {
				break
			}
		}
	}
	return
}

// MkParentDirs 为父路径，创建（一级或多级）目录.
/*
@param filePaths （文件 || 目录）路径s（相对路径 || 绝对路径）

e.g.
("")	=> nil
(".")	=> nil
*/
func MkParentDirs(paths ...string) (err error) {
	for _, path := range paths {
		// Richelieu: 为防止 import cycle，不直接使用 pathKit.GetParentDir()
		//dirPath := pathKit.GetParentDir(path)
		dirPath := filepath.Dir(path)

		err = MkDirs(dirPath)
		if err != nil {
			break
		}
	}
	return
}
