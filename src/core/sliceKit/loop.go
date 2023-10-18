package sliceKit

import "github.com/samber/lo"
import lop "github.com/samber/lo/parallel"

// ForEach 遍历（内部是for range，遍历前会拷贝一份s然后遍历它）.
/*
@param s		可以为nil
@param iteratee	不能为nil

e.g.
	s := []int{0, 1, 2}
	fmt.Println(s) // [0 1 2]

	sliceKit.ForEach(s, func(item int, index int) {
		s[index] = item + 1
	})
	fmt.Println(s) // [1 2 3]
*/
func ForEach[T any](s []T, iteratee func(item T, index int)) {
	lo.ForEach(s, iteratee)
}

// ForEachInParallel 遍历（内部是for range + goroutine，遍历前会拷贝一份s然后遍历它）.
/*
PS: 多协程，并发.

@param s		可以为nil
@param iteratee	不能为nil
*/
func ForEachInParallel[T any](s []T, iteratee func(item T, index int)) {
	lop.ForEach(s, iteratee)
}
