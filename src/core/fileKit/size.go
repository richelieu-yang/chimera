package fileKit

import (
	"os"
	"path/filepath"
)

// GetSize 获取文件（或目录）的大小.
func GetSize(path string) (int64, error) {
	if err := AssertExist(path); err != nil {
		return 0, err
	}

	if IsFile(path) {
		return GetFileSize(path)
	}
	return GetDirSize(path)
}

// GetFileSize 获取文件的大小.
func GetFileSize(filePath string) (int64, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return 0, err
	}

	info, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// GetDirSize 获取目录的大小（包含其内文件和目录）.
/*
参考:
golang获取文件/目录（包含下面的文件）的大小: https://blog.csdn.net/n_fly/article/details/117080173
*/
func GetDirSize(dirPath string) (int64, error) {
	if err := AssertExistAndIsDir(dirPath); err != nil {
		return 0, err
	}

	var bytes int64
	err := filepath.Walk(dirPath, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			bytes += info.Size()
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return bytes, nil
}
