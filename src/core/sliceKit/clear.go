//go:build !go1.21

package sliceKit

// Clear
/*
Deprecated: go <= 1.20，性能较差，还不如新建个空的然后覆盖.
*/
func Clear[T any](s []T) {
	var zero T
	for i, _ := range s {
		s[i] = zero
	}
}
