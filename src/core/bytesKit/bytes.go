package bytesKit

import "bytes"

// Compare 按照字典序比较两个字节切片的大小.
/*
@param a 可以为nil
@param b 可以为nil
@return -1:	a < b
		0:	a == b
		1:	a > b
*/
var Compare func(a, b []byte) int = bytes.Compare

// Equals
/*
@param a 可以为nil
@param b 可以为nil

e.g.
	var a = []byte("abcd")
	var b = []byte("abcd")
	println(bytesKit.Equals(a, b)) // true
	b = []byte("abcd1")
	println(bytesKit.Equals(a, b))     // false
	println(bytesKit.Equals(a, nil))   // false
	println(bytesKit.Equals(nil, nil)) // true
*/
func Equals(a, b []byte) bool {
	return bytes.Compare(a, b) == 0
}
