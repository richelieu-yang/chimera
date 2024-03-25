package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"path/filepath"
	"time"
)

var (
	// Chmod 使用指定的权限，更改指定路径的文件权限.
	Chmod func(path string, mode os.FileMode) (err error) = gfile.Chmod

	// CutAndPaste 剪贴.
	CutAndPaste func(src string, dst string) (err error) = gfile.Move

	// Remove 删除文件（或目录）.
	/*
		PS: 如果是目录且内部有文件或目录，也会一并删除.
	*/
	Remove func(path string) (err error) = gfile.Remove

	// Delete 删除文件（或目录）.
	Delete func(path string) (err error) = Remove

	// Truncate 裁剪文件为指定大小.
	/*
		PS:
		(1) 如果给定文件路径是软链，将会修改源文件;
		(2) If there is an error, it will be of type *PathError.

		@param size 如果为0，则清空文件内容
	*/
	Truncate func(path string, size int) (err error) = gfile.Truncate
)

// EmptyDir 清空目录：删掉目录中的文件和子目录（递归），但该目录本身不会被删掉.
/*
@param dirPath 可以不存在（此时将返回nil）
*/
func EmptyDir(dirPath string) error {
	if !Exists(dirPath) {
		return nil
	}
	if err := AssertExistAndIsDir(dirPath); err != nil {
		return err
	}

	// 遍历目录
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, dirEntry := range dirEntries {
		path := filepath.Join(dirPath, dirEntry.Name())
		if err := Remove(path); err != nil {
			return err
		}
	}
	return nil
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
