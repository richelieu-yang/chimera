package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

var (
	// IsReadable 是否有 读 权限?
	/*
		e.g.
			("") => false
	*/
	IsReadable func(path string) bool = gfile.IsReadable

	// IsWritable 是否有 写 权限?
	/*
		e.g.
			("") => false
	*/
	IsWritable func(path string) bool = gfile.IsWritable
)

// GetFileMode get mode and permission bits of file/directory
func GetFileMode(path string) (os.FileMode, error) {
	if err := AssertExist(path); err != nil {
		return 0, err
	}

	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Mode(), nil
}
