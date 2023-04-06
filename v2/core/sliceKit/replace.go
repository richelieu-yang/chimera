package sliceKit

import "github.com/samber/lo"

// Replace
/*
@param collection 	可以为nil
@param n			替换的个数（如果有这么多的话）
					(1) > 0，替换指定个数
					(2) = 0，一个都不替换
					(3) < 0，替换所有
@return 保底为length==0的切片实例

e.g.
# 替换0个
([]int{0, 1, 2, 2, 3, 3, 3}, 2, 9, 0) 	=> [0 1 2 2 3 3 3]
# 替换2个
([]int{0, 1, 2, 2, 3, 3, 3}, 3, 9, 2) 	=> [0 1 2 2 9 9 3]
# 替换所有
([]int{0, 1, 2, 2, 3, 3, 3}, 3, 9, -1) 	=> [0 1 2 2 9 9 9]
*/
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	return lo.Replace(collection, old, new, n)
}

// ReplaceAll 替换所有.
/*
@param collection 	可以为nil
@return 保底为length==0的切片实例
*/
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return lo.ReplaceAll(collection, old, new)
}
