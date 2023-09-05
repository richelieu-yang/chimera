package compareKit

import "github.com/google/go-cmp/cmp"

// Equal 是否相等？
/*
PS:
(1) 结果传参分别为 结构体实例 和 结构体实例指针，将返回false.
*/
var Equal func(x, y interface{}, opts ...cmp.Option) bool = cmp.Equal

// Diff 获取差异.
var Diff func(x, y interface{}, opts ...cmp.Option) string = cmp.Diff
