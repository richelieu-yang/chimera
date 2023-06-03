package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"time"
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
)

// GetModificationTime 获取文件（或目录）的修改时间
/*
@param path 传参""将返回err（Stat : The system cannot find the path specified.）
*/
func GetModificationTime(path string) (time.Time, error) {
	info, err := Stat(path)
	if err != nil {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

// SetModificationTime 修改文件（或目录）的修改时间
/*
PS:
(1) 也会同时修改文件（或目录）的访问时间；
(2) 修改目录的修改时间，将不会影响该目录下的文件或目录；
(3) 传参t可以晚于当前时间.

@param path 传参""将返回error（chtimes : The system cannot find the path specified.）
*/
func SetModificationTime(path string, t time.Time) error {
	return os.Chtimes(path, t, t)
}
