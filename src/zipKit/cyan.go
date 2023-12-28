package zipKit

import "github.com/dablelv/cyan/zip"

var (
	// Zip 压缩文件（或目录）
	/*
	   PS:
	   (1) 缺点: 第三方库dablelv/go-huge-util暂不支持带密码的压缩.
	   (2) 如果 zipPath 对应的是一个已经存在的文件，将会"覆盖"该文件的内容；
	   	如果 zipPath 对应的是一个已经存在的目录，将会返回error（open {path}: is a directory）.

	   @param zipPath 	e.g."archive.zip"
	   @param paths	多个文件（或目录）的路径
	*/
	Zip func(zipPath string, paths ...string) error = zip.Zip

	// Unzip 解压
	/*
	   @param zipPath e.g."archive.zip"
	*/
	Unzip func(zipPath, dstDir string) error = zip.Unzip
)
