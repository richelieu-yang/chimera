package fileKit

import (
	"os"
	"time"
)

// Exist 判断文件（或目录）是否存在.
/*
@param path 绝对路径 || 相对路径

e.g.
("") => false
(" ") => false
*/
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func NotExist(path string) bool {
	return !Exist(path)
}

// IsDir
/*
@return 如果传参path对应的文件或目录不存在，将返回false
*/
func IsDir(path string) bool {
	if fileInfo, err := os.Stat(path); err != nil {
		return false
	} else {
		return fileInfo.IsDir()
	}
}

// IsFile
/*
@return 如果传参path对应的文件或目录不存在，将返回false
*/
func IsFile(path string) bool {
	if fileInfo, err := os.Stat(path); err != nil {
		return false
	} else {
		return !fileInfo.IsDir()
	}
}

// GetFileInfo 获取文件（或目录）信息
func GetFileInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// GetModificationTime 获取文件（或目录）的修改时间
/*
@param path 传参""将返回err（Stat : The system cannot find the path specified.）
*/
func GetModificationTime(path string) (time.Time, error) {
	info, err := os.Stat(path)
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
