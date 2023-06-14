package bytesKit

import "bytes"

var Compare func(a, b []byte) int = bytes.Compare

// Equals
/*
e.g.
	var a = []byte("abcd")
	var b = []byte("abcd")
	println(bytesKit.Equals(a, b)) // true
	b = []byte("abcd1")
	println(bytesKit.Equals(a, b)) // false
*/
func Equals(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}
