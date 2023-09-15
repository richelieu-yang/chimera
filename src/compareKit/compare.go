package compareKit

import "github.com/google/go-cmp/cmp"

// Equal 是否相等？
/*
!!!:
(1) 如果传参结构体（或其内部的结构体）实现了 "(T) Equal(T) bool" 或者 "(T) Equal(I) bool"，
	方法的receiver必须为 "值类型"，这样的话，无论比较 结构体实例 还是 结构体实例指针 都可以.

PS:
(1) 如果传参分别为 结构体实例 和 结构体实例指针，将返回false.
(2) 传参x、y值都为nil的情况下，	(a) 类型不同，将返回 false;
							(b) 类型相同，将返回 true.
*/
var Equal func(x, y interface{}, opts ...cmp.Option) bool = cmp.Equal

// Diff 获取差异.
/*
@return 如果为""则说明两个传参一致.
*/
var Diff func(x, y interface{}, opts ...cmp.Option) string = cmp.Diff
