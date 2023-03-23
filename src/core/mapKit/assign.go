package mapKit

import (
	"github.com/samber/lo"
)

// Merge 合并多个map实例（合并到最左边的map实例中）
/*
PS: 存在key相同的情况，以最右边的为准.

@param maps 可以为nil; 内部可以有nil
@return 保底为空的map实例（不为nil）

e.g.
	m := mapKit.Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
	)
	fmt.Println(m) // map[a:1 b:3 c:4]

e.g.1
	m := mapKit.Merge[string, int](
		map[string]int{"a": 1, "b": 2},
		map[string]int{"b": 3, "c": 4},
		map[string]int{"b": 5},
	)
	fmt.Println(m) // map[a:1 b:5 c:4]
*/
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	return lo.Assign(maps...)
}
