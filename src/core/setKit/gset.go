package setKit

import "github.com/gogf/gf/v2/container/gset"

// NewGset Deprecated: 存储的元素类型为 interface{}.
/*
@param safe 是否并发安全？
*/
var NewGset func(safe ...bool) *gset.Set = gset.NewSet

// NewGsetFrom Deprecated: 存储的元素类型为 interface{}.
/*
@param items 	(1) 切片实例（元素类型不限）
				(2) 可以为nil
@param safe 	是否并发安全？
*/
var NewGsetFrom func(items interface{}, safe ...bool) *gset.Set = gset.NewFrom
