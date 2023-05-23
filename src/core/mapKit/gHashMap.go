package mapKit

import "github.com/gogf/gf/v2/container/gmap"

func NewAnyAnyMap(safe ...bool) *gmap.AnyAnyMap {
	return gmap.NewAnyAnyMap(safe...)
}

func NewAnyAnyMapFrom(data map[interface{}]interface{}, safe ...bool) *gmap.AnyAnyMap {
	return gmap.NewAnyAnyMapFrom(data, safe...)
}

func NewIntAnyMap(safe ...bool) *gmap.IntAnyMap {
	return gmap.NewIntAnyMap(safe...)
}

func NewIntAnyMapFrom(data map[int]interface{}, safe ...bool) *gmap.IntAnyMap {
	return gmap.NewIntAnyMapFrom(data, safe...)
}

func NewStrAnyMap(safe ...bool) *gmap.StrAnyMap {
	return gmap.NewStrAnyMap(safe...)
}

func NewStrAnyMapFrom(data map[string]interface{}, safe ...bool) *gmap.StrAnyMap {
	return gmap.NewStrAnyMapFrom(data, safe...)
}

func NewIntIntMap(safe ...bool) *gmap.IntIntMap {
	return gmap.NewIntIntMap(safe...)
}

func NewIntIntMapFrom(data map[int]int, safe ...bool) *gmap.IntIntMap {
	return gmap.NewIntIntMapFrom(data, safe...)
}

func NewStrStrMap(safe ...bool) *gmap.StrStrMap {
	return gmap.NewStrStrMap(safe...)
}

func NewStrStrMapFrom(data map[string]string, safe ...bool) *gmap.StrStrMap {
	return gmap.NewStrStrMapFrom(data, safe...)
}

func NewIntStrMap(safe ...bool) *gmap.IntStrMap {
	return gmap.NewIntStrMap(safe...)
}

func NewIntStrMapFrom(data map[int]string, safe ...bool) *gmap.IntStrMap {
	return gmap.NewIntStrMapFrom(data, safe...)
}

func NewStrIntMap(safe ...bool) *gmap.StrIntMap {
	return gmap.NewStrIntMap(safe...)
}

func NewStrIntMapFrom(data map[string]int, safe ...bool) *gmap.StrIntMap {
	return gmap.NewStrIntMapFrom(data, safe...)
}
