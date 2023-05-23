package mapKit

import "github.com/gogf/gf/v2/container/gmap"

func NewListMap(safe ...bool) *gmap.ListMap {
	return gmap.NewListMap(safe...)
}

func NewListMapFrom(data map[interface{}]interface{}, safe ...bool) *gmap.ListMap {
	return gmap.NewListMapFrom(data, safe...)
}
