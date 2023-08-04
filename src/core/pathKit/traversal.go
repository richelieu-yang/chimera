package pathKit

import (
	"os"
	"path/filepath"
)

// Walk 遍历目录（包含传参root）.
var Walk func(root string, fn filepath.WalkFunc) error = filepath.Walk

// ReadDir 遍历目录（不包含传参root）.
var ReadDir func(name string) ([]os.DirEntry, error) = os.ReadDir
