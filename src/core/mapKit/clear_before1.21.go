//go:build !go1.21

package mapKit

// Clear
/*
Deprecated: go <= 1.20的情况下，性能较差，还不如新建个 空的map实例 然后覆盖.
*/
func Clear[K comparable, V any](m map[K]V) {
	for k, _ := range m {
		delete(m, k)
	}
}
