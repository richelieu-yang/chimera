//go:build go1.21

package mapKit

// Clear 清空所有map的键值对，clear后，我们将得到一个empty map
func Clear[K comparable, V any](m map[K]V) {
	clear(m)
}
