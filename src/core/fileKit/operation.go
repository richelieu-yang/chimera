package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
	"path/filepath"
	"time"
)

var (
	// Chmod 修改权限
	Chmod func(path string, mode os.FileMode) (err error) = gfile.Chmod

	Move   func(src string, dst string) (err error) = gfile.Move
	Rename func(src string, dst string) (err error) = gfile.Move

	// Remove 删除文件（或目录）
	/*
		PS: 如果是目录且内部有文件或目录，也会一并删除.
	*/
	Remove = gfile.Remove
)

// NewFile 创建文件（读写权限、文件不存在就创建、打开并清空文件）.
func NewFile(filePath string) (*os.File, error) {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	// flag: O_RDWR|O_CREATE|O_TRUNC
	return os.Create(filePath)
}

// NewFileInAppendMode 创建文件（读写权限、文件不存在就创建、追加模式）.
func NewFileInAppendMode(filePath string) (*os.File, error) {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
}

// NewTemporaryFile 在指定目录下，生成临时文件.
/*
@param dirPath 如果为""，临时文件将生成在 系统临时目录 内；如果为"."，临时文件将生成在 当前目录 内.

e.g.
pattern: "tempfile_test" 		=> 临时文件的文件名: "tempfile_test2594316144"
pattern: "tempfile_test*" 		=> 临时文件的文件名: "tempfile_test827818253"
pattern: "tempfile_test*.xyz" 	=> 临时文件的文件名: "tempfile_test3617672388.xyz"
*/
func NewTemporaryFile(dirPath, pattern string) (*os.File, error) {
	if err := AssertNotExistOrIsDir(dirPath); err != nil {
		return nil, err
	}
	if err := MkDirs(dirPath); err != nil {
		return nil, err
	}

	if err := strKit.AssertNotBlank(pattern, "pattern"); err != nil {
		return nil, err
	}

	return os.CreateTemp(dirPath, pattern)
}

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
