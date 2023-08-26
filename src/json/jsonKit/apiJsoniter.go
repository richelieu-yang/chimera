//go:build !amd64 || !go1.16

package jsonKit

func init() {
	useJsonIterator()
}
