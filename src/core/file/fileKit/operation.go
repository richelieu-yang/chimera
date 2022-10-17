package fileKit

import (
	"io"
	"os"
	"path/filepath"
)

// NewTemporaryFile 在指定目录下，生成临时文件.
/*
@param dirPath 如果为""，临时文件将生成在 系统临时目录 内；如果为"."，临时文件将生成在 当前目录 内.

e.g.
pattern: "tempfile_test" 		=> 临时文件的文件名: "tempfile_test2594316144"
pattern: "tempfile_test*" 		=> 临时文件的文件名: "tempfile_test827818253"
pattern: "tempfile_test*.xyz" 	=> 临时文件的文件名: "tempfile_test3617672388.xyz"
*/
func NewTemporaryFile(dirPath, pattern string) (*os.File, error) {
	// check dirPath
	if err := AssertNotExistOrIsDir(dirPath); err != nil {
		return nil, err
	}
	if err := MkDirs(dirPath); err != nil {
		return nil, err
	}

	return os.CreateTemp(dirPath, pattern)
}

// NewFile 创建文件.
/*
PS: 如果文件已经存在，会覆盖掉它.
*/
func NewFile(filePath string) (*os.File, error) {
	// 检查 filePath
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return os.Create(filePath)
}

// WriteToFile 将数据（字节流）写到文件中.
/*
@param target 目标文件的路径（不存在的话，会创建一个新的文件；存在且是个文件的话，会覆盖掉旧的（并不会加到该文件的最后面））
*/
func WriteToFile(data []byte, dest string) error {
	// 检查 dest
	if err := AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := MkParentDirs(dest); err != nil {
		return err
	}

	return os.WriteFile(dest, data, os.ModePerm)
}

// CopyFile 复制单个文件.
/*
@param src	一个已经存在的文件
@param dest	一个已经存在的文件（会覆盖） || 一个不存在的文件
@return 第一个返回值: the number of bytes copied（单位为byte）
*/
func CopyFile(src, dest string) (int64, error) {
	// 检查 src
	if err := AssertExistAndIsFile(src); err != nil {
		return 0, err
	}
	// 检查 dest
	if err := AssertNotExistOrIsFile(dest); err != nil {
		return 0, err
	}
	if err := MkParentDirs(dest); err != nil {
		return 0, err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	return io.Copy(destFile, srcFile)
}

// CopyDir 将 src目录的所有内容 复制到 dest目录 中.
/*
PS:
(1) src目录下如果还有目录，会递归（空目录也会复制过去）；
(2) 类似于Linux的 cp -r 命令.

@param src	一个已经存在的目录
@param dest	一个已经存在的目录 || 一个不存在的目录
*/
func CopyDir(src, dest string) error {
	// 检查 src
	if err := AssertExistAndIsDir(src); err != nil {
		return err
	}
	// 检查 dest
	if err := AssertNotExistOrIsDir(dest); err != nil {
		return err
	}
	if err := MkDirs(dest); err != nil {
		return err
	}

	// 遍历
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		// 此次跳过
		if relPath == "." {
			return nil
		}
		tmpSrc := filepath.Join(src, relPath)
		tmpDest := filepath.Join(dest, relPath)
		if info.IsDir() {
			return CopyDir(tmpSrc, tmpDest)
		}
		_, err = CopyFile(tmpSrc, tmpDest)
		return err
	})
}

// Delete 删除文件或目录（内部有文件或目录也会一并删除）.
/*
PS:
(1) 传参""，不会报错；
(2) path对应的文件或目录不存在，不会报错；
(3) Windows，若文件被锁定，会报错：panic: remove xxx(path): The process cannot access the file because it is being used by another process.
(4) Windows，若目录中存在被锁定的文件，能删的会被删掉，然后报错：panic: remove xxx(path): The process cannot access the file because it is being used by another process.
*/
func Delete(path string) error {
	return os.RemoveAll(path)
}

// EmptyDir 清空目录：删掉目录中的文件和子目录（递归），但该目录本身不会被删掉.
func EmptyDir(dirPath string) error {
	if NotExist(dirPath) {
		return nil
	}
	if err := AssertExistAndIsDir(dirPath); err != nil {
		return err
	}

	// 遍历目录
	dir, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	// 在删的过程中，就算报错了（e.g.文件被占用），也要全部都删一遍，返回第一个error
	for _, f := range dir {
		err1 := os.RemoveAll(filepath.Join(dirPath, f.Name()))
		if err1 != nil && err == nil {
			err = err1
		}
	}
	return err
}
