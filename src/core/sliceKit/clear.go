//go:build !go1.21

package sliceKit

// Clear 保持slice的长度（len）和容量（cap），但将所有slice内已存在的元素(len个)都置为元素类型的零值.
/*
Deprecated: go <= 1.20，性能较差，还不如新建个空的然后覆盖.
*/
func Clear[T any](s []T) {
	var zero T
	for i, _ := range s {
		s[i] = zero
	}
}
