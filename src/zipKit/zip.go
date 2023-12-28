package zipKit

import "github.com/duke-git/lancet/v2/fileutil"

var (
	// Zip zip压缩文件.
	/*
		@param path 可以是 文件 || 目录
	*/
	Zip func(path string, destPath string) error = fileutil.Zip

	// UnZip zip解压缩文件并保存在目录中.
	UnZip func(zipFile string, destPath string) error = fileutil.UnZip

	// ZipAppendEntry 通过将单个文件或目录追加到现有的zip文件.
	ZipAppendEntry func(fpath string, destPath string) error = fileutil.ZipAppendEntry

	// IsZipFile 判断文件是否是zip压缩文件.
	IsZipFile func(filepath string) bool = fileutil.IsZipFile
)
