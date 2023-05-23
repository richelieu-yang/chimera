package mapKit

import (
	"github.com/gogf/gf/v2/container/gmap"
)

// NewTreeMap
/*
PS: 当需要让返回结果按照自然升序排列时使用TreeMap.

@param comparator e.g. gutil.ComparatorInt
*/
func NewTreeMap(comparator func(v1, v2 interface{}) int, safe ...bool) *gmap.TreeMap {
	return gmap.NewTreeMap(comparator, safe...)
}

func NewTreeMapFrom(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *gmap.TreeMap {
	return gmap.NewTreeMapFrom(comparator, data, safe...)
}
