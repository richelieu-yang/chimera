//go:build go1.21

package sliceKit

// Clear 保持slice的长度（len）和容量（cap），但将所有slice内已存在的元素(len个)都置为元素类型的零值.
func Clear[T any](s []T) {
	clear(s)
}
