package sliceKit

import "github.com/richelieu42/chimera/v2/src/copyKit"

// Copy 浅克隆（浅拷贝）
/*
参考:
golang复制切片的方法（避免操作同一底层数组） https://blog.csdn.net/weixin_43970884/article/details/126051345

@param src 	可以为nil
@return 保底为len==0的slice实例

e.g.	浅拷贝
	s0 := []string{"0", "1", "2"}
	s1 := sliceKit.Copy(s0)

	s1[0] = "3"
	fmt.Println(s0) // [0 1 2]
	fmt.Println(s1) // [3 1 2]

e.g.1	传参为nil
	s := sliceKit.Copy([]int(nil))
	fmt.Println(s)        // []
	fmt.Println(len(s))   // 0
	fmt.Println(s != nil) // true
*/
func Copy[T any](src []T) []T {
	dolly := make([]T, len(src))
	copy(dolly, src)
	return dolly
}

// Copy1 浅克隆（浅拷贝）
/*
Deprecated: 并不是我想要的"浅拷贝".

e.g.
	s0 := []string{"0", "1", "2"}
	s1 := sliceKit.Copy1(s0)

	s1[0] = "3"
	fmt.Println(s0) // [3 1 2]
	fmt.Println(s1) // [3 1 2]
*/
func Copy1[T any](src []T) []T {
	return append(src)
}

// Copy2 浅克隆（浅拷贝）
/*
Deprecated: 并不是我想要的"浅拷贝".

e.g.
	s0 := []string{"0", "1", "2"}
	s1 := sliceKit.Copy2(s0)

	s1[0] = "3"
	fmt.Println(s0) // [3 1 2]
	fmt.Println(s1) // [3 1 2]
*/
func Copy2[T any](src []T) []T {
	return src[:]
}

// DeepCopy 深拷贝
/*
e.g.
	([]string(nil)) => (nil, nil)

e.g.1
	s := []*Bean{{Id: 0}, {Id: 1}}
	s1, err := sliceKit.DeepCopy(s)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%p\n", &s)   // 0x1400000c048
	fmt.Printf("%p\n", s[0]) // 0x140000220a0
	fmt.Printf("%p\n", s[1]) // 0x140000220a8

	fmt.Println("==================")

	fmt.Printf("%p\n", &s1)   // 0x1400000c060
	fmt.Printf("%p\n", s1[0]) // 0x140000220b0
	fmt.Printf("%p\n", s1[1]) // 0x140000220d8
*/
func DeepCopy[T any](src []T) ([]T, error) {
	return copyKit.DeepCopy(src)
}

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
