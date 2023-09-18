package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

var (
	// Open （只读模式）打开文件/目录.
	/*
		!!!: 只读权限 打开.
	*/
	Open func(path string) (*os.File, error) = gfile.Open

	// OpenFile （以指定 flag 和 perm）打开文件/目录.
	/*
		@param flag 详见".info"
		@param perm	(1) 可以参考 "fileKit/consts.go"
					(2) e.g.0666 || os.ModePerm ...
	*/
	OpenFile func(path string, flag int, perm os.FileMode) (*os.File, error) = gfile.OpenFile

	// Create 创建文件（目录不行）.
	/*
		PS:
		(1) flag == O_RDWR|O_CREATE|O_TRUNC
		(2) 读写权限（0666）；
		(3) path不存在，会创建；
		(4) path存在 && 是文件，会清空该文件内容；
		(5) path存在 && 是目录，会返回error.
	*/
	Create func(path string) (*os.File, error) = gfile.Create
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
