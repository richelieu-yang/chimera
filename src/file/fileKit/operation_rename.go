package fileKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"os"
)

// Move 移动.
var Move func(src, dst string) error = Rename

// Rename 重命名.
/*
PS: 参考 gfile.Rename.
*/
func Rename(src, dst string) error {
	if err := AssertExist(src); err != nil {
		return err
	}
	if err := MkParentDirs(dst); err != nil {
		return err
	}

	err := os.Rename(src, dst)
	if err != nil {
		err = errorKit.Wrapf(err, `fail to rename from "%s" to "%s"`, src, dst)
	}
	return err
}

// RenameInSameDir 同目录下重命名.
func RenameInSameDir(path string, name string) error {
	if err := AssertExist(path); err != nil {
		return err
	}
	if err := strKit.AssertNotBlank(name, "name"); err != nil {
		return err
	}

	parentDir := pathKit.ParentDir(path)
	return Rename(path, pathKit.Join(parentDir, name))
}
