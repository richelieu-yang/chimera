//go:build !go1.21

package mapKit

func Clear[K comparable, V any](m map[K]V) {
	for k, _ := range m {
		delete(m, k)
	}
}
