//go:build go1.21

package sliceKit

func Clear[T any](s []T) {
	clear(s)
}
