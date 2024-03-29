package sliceKit

import "github.com/samber/lo"

// Drop 从"左"开始，丢弃n个数据
/*
@param collection	(1) 可以为nil
					(2) 不会修改传参collection
@param n			丢弃数据的个数
@return 非nil的slice实例（len >= 0）

e.g.	不会修改传参collection
	s := []int{0, 1, 2, 3, 4, 5}
	s1 := sliceKit.Drop(s, 2)
	fmt.Println(s)  // [0 1 2 3 4 5]
	fmt.Println(s1) // [2 3 4 5]

	s1[0] = 9
	fmt.Println(s)  // [0 1 2 3 4 5]
	fmt.Println(s1) // [9 3 4 5]
*/
func Drop[T any](collection []T, n int) []T {
	return lo.Drop(collection, n)
}

// DropRight 从"右"开始，丢弃n个数据
/*
e.g.
([]int{0, 1, 2, 3, 4, 5}, 2) => []int{0, 1, 2, 3}
*/
func DropRight[T any](collection []T, n int) []T {
	return lo.DropRight(collection, n)
}

// DropWhile
/*
@param collection (1)可以为nil (2)此传参不会被修改
@return 不为nil，len >= 0

e.g.
	s := []string{"a", "aa", "aaa", "aa", "aa"}
	s1 := sliceKit.DropWhile(s, func(val string) bool {
		return len(val) <= 2
	})
	fmt.Println(s)  // [a aa aaa aa aa]
	fmt.Println(s1) // [aaa aa aa]
*/
func DropWhile[T any](collection []T, predicate func(item T) bool) []T {
	return lo.DropWhile(collection, predicate)
}

// DropRightWhile
/*
@param collection (1)可以为nil (2)此传参不会被修改
@return 不为nil，len >= 0

e.g.
	s := []string{"a", "aa", "aaa", "aa", "aa"}
	s1 := sliceKit.DropRightWhile(s, func(val string) bool {
		return len(val) <= 2
	})
	fmt.Println(s)  // [a aa aaa aa aa]
	fmt.Println(s1) // [aaa aa aa]
*/
func DropRightWhile[T any](collection []T, predicate func(item T) bool) []T {
	return lo.DropRightWhile(collection, predicate)
}
