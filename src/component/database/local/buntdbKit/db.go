package buntdbKit

import (
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/tidwall/buntdb"
)

const (
	MemoryPath = ":memory:"
)

// Open
/*
@param 存储数据的文件路径（会自动创建父目录）
*/
func Open(path string) (*buntdb.DB, error) {
	if path == MemoryPath {
		return OpenInMemory()
	}

	if err := fileKit.AssertNotExistOrIsFile(path); err != nil {
		return nil, err
	}
	return buntdb.Open(path)
}

// OpenInMemory 不会持久化到磁盘上.
func OpenInMemory() (*buntdb.DB, error) {
	return buntdb.Open(MemoryPath)
}
