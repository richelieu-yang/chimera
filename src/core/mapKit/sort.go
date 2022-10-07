package mapKit

import (
	"gitee.com/richelieu042/go-scales/src/core/sliceKit"
	"golang.org/x/exp/constraints"
)

// Sort 排序map实例（根据key升序||降序）
/*
Deprecated: 没有什么实际的使用场景.

@param m	可以为nil
@param args	true: 降序；false: 升序（默认）
@return 	可能为nil
*/
func Sort[K constraints.Ordered, V any](m map[K]V, args ...bool) map[K]V {
	if len(m) <= 1 {
		return m
	}

	rst := make(map[K]V)
	keys := GetKeySlice(m)
	keys = sliceKit.Sort(keys, args...)
	for _, key := range keys {
		rst[key] = m[key]
	}
	return rst
}
