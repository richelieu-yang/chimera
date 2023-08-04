package pathKit

import "path/filepath"

// Walk 遍历目录.
/*
PS: 详见"Golang.wps".
*/
var Walk func(root string, fn filepath.WalkFunc) error = filepath.Walk
