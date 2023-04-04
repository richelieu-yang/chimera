package interfaceKit

import "github.com/samber/lo"

// GetZeroValue 获取类型T的零值.
func GetZeroValue[T any]() T {
	return lo.Empty[T]()
}

// IsZeroValue 传参v是否是类型T的零值？
/*
e.g.
[interface{}](nil) => true
*/
func IsZeroValue[T comparable](v T) bool {
	return lo.IsEmpty(v)
}
