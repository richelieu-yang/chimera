package sliceKit

// Copy 将src的元素复制到dest中（可能会修改dest的内容，但dest的内存地址不变！！！）
/*
PS:
(1) src和dest中只要有一个为nil（包括两个都为nil的情况），返回值必定为0 && 不会修改dest；
(2) src和dest都不为nil的情况下，len(src) <= len(dest)，复制 src的前n个元素 并用他们覆盖 dest的前n个元素（n == len(src)，可能为0）；
(3) src和dest都不为nil的情况下，len(src) > len(dest)，复制 src的前n个元素 并用他们覆盖 dest的前n个元素（n == len(dest)，可能为0）.

@param src	可以为nil
@param dest	可以为nil
@return 	被复制的元素的个数（即len(src)和len(dest)间的最小值；>= 0）

e.g. 	nil的情况
[byte](nil, nil) 				=> 0
[byte]([]byte("abc"), nil) 		=> 0
[byte](nil, []byte("abcde")) 	=> 0
e.g.1	正常情况
([]byte("012"), []byte("abcde")) => 3（dest变为[]byte("012de")）
([]byte("01234"), []byte("abc")) => 3（dest变为[]byte("012")）
*/
func Copy[T any](src, dest []T) int {
	return copy(dest, src)
}
