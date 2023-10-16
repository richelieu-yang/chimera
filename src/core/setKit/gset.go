package setKit

import "github.com/gogf/gf/v2/container/gset"

// NewGset
/*
@param safe 是否并发安全？
*/
var NewGset func(safe ...bool) *gset.Set = gset.NewSet

// NewGsetFrom
/*
@param items 	(1) 切片实例（元素类型不限）
				(2) 可以为nil
@param safe 	是否并发安全？
*/
var NewGsetFrom func(items interface{}, safe ...bool) *gset.Set = gset.NewFrom
