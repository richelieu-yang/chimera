package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"path/filepath"
)

var (
	Exists func(path string) bool = gfile.Exists
	IsFile func(path string) bool = gfile.IsFile
	IsDir  func(path string) bool = gfile.IsDir

	// Stat 获取文件（或目录）信息
	Stat func(path string) (os.FileInfo, error) = gfile.Stat

	IsReadable func(path string) bool = gfile.IsReadable
	IsWritable func(path string) bool = gfile.IsWritable

	// IsEmpty checks whether the given `path` is empty.
	// If `path` is a folder, it checks if there's any file under it.
	// If `path` is a file, it checks if the file size is zero.
	//
	// Note that it returns true if `path` does not exist.
	IsEmpty func(path string) bool = gfile.IsEmpty

	// GetBaseName 获取 完整的文件名.
	/*
		e.g.
		/var/www/file.js -> file.js
		file.js          -> file.js
	*/
	GetBaseName func(path string) string = gfile.Basename

	// GetName 获取 前缀.
	/*
		e.g.
		/var/www/file.js -> file
		file.js          -> file
	*/
	GetName func(path string) string = gfile.Name

	// GetExt 获取 后缀（带"."）
	/*
		e.g.
		main.go  => .go
		api.json => .json
	*/
	GetExt func(path string) string = gfile.Ext

	// GetExtName 获取 后缀（不带"."）
	/*
		e.g.
		main.go  => go
		api.json => json
	*/
	GetExtName func(path string) string = gfile.ExtName
)

// GetSize 获取文件（或目录）的大小.
func GetSize(path string) (int64, error) {
	if err := AssertExist(path); err != nil {
		return 0, err
	}

	if IsFile(path) {
		return getFileSize(path)
	}
	return getDirSize(path)
}

// getFileSize 获取文件的大小.
func getFileSize(filePath string) (int64, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// getDirSize 获取目录的大小（包含其内文件和目录）.
/*
参考:
golang获取文件/目录（包含下面的文件）的大小: https://blog.csdn.net/n_fly/article/details/117080173
*/
func getDirSize(dirPath string) (int64, error) {
	var bytes int64
	err := filepath.Walk(dirPath, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			bytes += info.Size()
		}
		// 如果 err != nil，将中止遍历
		return err
	})
	if err != nil {
		return 0, err
	}
	return bytes, nil
}
