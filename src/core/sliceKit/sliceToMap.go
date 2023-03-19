package sliceKit

import "github.com/samber/lo"

// SliceToMap slice实例 => map实例
/*
@param s			可以为nil
@param transform	不能为nil（除非s == nil），否则会导致panic: runtime error: invalid memory address or nil pointer dereference
@return 			必定不为nil（保底为空的map实例）

e.g.
	m := sliceKit.SliceToMap[string, string, string]([]string{"0", "1"}, func(item string) (string, string) {
		return "key" + item, "value" + item
	})
	fmt.Println(m)	// map[key0:value0 key1:value1]
*/
func SliceToMap[T any, K comparable, V any](s []T, transform func(item T) (K, V)) map[K]V {
	return lo.Associate(s, transform)
}
