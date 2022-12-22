package sliceKit

import "github.com/richelieu42/go-scales/src/copyKit"

// CopyToDest 将src的元素复制到dest中（可能会修改dest的内容，但dest的内存地址不变！！！）
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
func CopyToDest[T any](src, dest []T) int {
	return copy(dest, src)
}

// DeepCopy 深拷贝
/*
@return e.g. ([]string(nil)) => (nil, nil)
*/
func DeepCopy[T any](src []T) ([]T, error) {
	return copyKit.DeepCopy(src)
}

// Copy 浅克隆（浅拷贝）
/*
参考:
golang复制切片的方法（避免操作同一底层数组） https://blog.csdn.net/weixin_43970884/article/details/126051345

@param src 	可以为nil，将返回一个非nil的len==0的slice
@return		必定不为nil
*/
func Copy[T any](src []T) []T {
	if src == nil {
		return nil
	}

	dolly := make([]T, len(src))
	copy(dolly, src)
	return dolly
}

// Copy1 浅克隆（浅拷贝）
/*
Deprecated: 性能不如 Copy()

@param src 可以为nil，将返回nil
*/
func Copy1[T any](src []T) []T {
	return append(src)
}

// Copy2 浅克隆（浅拷贝）
/*
Deprecated: 返回值可能为nil

@param src 可以为nil，将返回nil
*/
func Copy2[T any](src []T) []T {
	return src[:]
}
