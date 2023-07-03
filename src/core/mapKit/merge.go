package mapKit

import (
	"github.com/samber/lo"
)

// Merge 合并多个map实例
/*
PS: 存在key相同的情况，以最右边的为准.

@param maps (1)可以为nil
			(2)内部可以有nil（甚至全是nil）
			(3)不会修改其中任何一个map实例
@return (1)保底为空的map实例（不为nil）
		(2)返回的是一个新的map实例

e.g.
	m := mapKit.Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"a": 10},
		map[string]int{"a": 100},
	)
	fmt.Println(m) // map[a:100 b:2]
*/
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {

	lo.Invert()

	return lo.Assign(maps...)
}
