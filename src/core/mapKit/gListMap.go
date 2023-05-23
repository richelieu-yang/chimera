package mapKit

import "github.com/gogf/gf/v2/container/gmap"

// NewListMap
/*
使用场景: 当需要按输入顺序返回结果时使用ListMap.
*/
func NewListMap(safe ...bool) *gmap.ListMap {
	return gmap.NewListMap(safe...)
}

func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *gmap.ListMap {
	return gmap.NewListMapFrom(data, safe...)
}
