package sliceKit

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
	"sort"
)

// SimpleSort （升序）排序
/*
@param s 可以为nil

e.g.
	s := []int{9, 0, -1, 1}
	sliceKit.SimpleSort(s)
	fmt.Println(s) // [-1 0 1 9]
*/
func SimpleSort[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

// SimpleSortByDesc （降序）排序
/*
@param s 可以为nil

e.g.
	s := []int{9, 0, -1, 1}
	sliceKit.SimpleSortByDesc(s)
	fmt.Println(s) // [9 1 0 -1]
*/
func SimpleSortByDesc[T constraints.Ordered](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}

// Sort 排序复杂类型的slice实例（元素的类型可以为自定义结构体）
/*
PS:
(1) 需要注意的是 sort.Sort() 函数不能保证数据排序是稳定的，
	如果需要保证数据排序稳定，可以使用 sort.Stable() 函数。（“稳定”的含义是原始数据中 a 和 b 的值相等，排序前 a 排在 b 的前面，排序后 a 仍排在 b 的前面）
(2) 现成的sort.Interface接口的实现: sort.IntSlice、sort.Float64Slice、sort.StringSlice.

@param data	可以为nil
*/
func Sort(data sort.Interface) {
	sort.Sort(data)
}

// SortStably 类似于 Sort，区别: SortStably 能保证数据排序是稳定的.
func SortStably(data sort.Interface) {
	sort.Stable(data)
}

// IsSorted 传参切片实例是否有序？
/*
@param s 	(1) 可以为nil
			(2) 如果 len(s) <= 1 ，返回值一定为true

e.g.
[int](nil)				=> true
([]string{})			=> true
([]string{"b"})			=> true
([]string{"b", "a"})	=> false
([]int{0, 1, 9, 100})	=> true
*/
func IsSorted[T constraints.Ordered](s []T) bool {
	return lo.IsSorted(s)
}

// IsSortedByKey 传参切片实例是否有序？
/*
@param s 	(1) 可以为nil
			(2) 如果 len(s) <= 1 ，返回值一定为true

e.g.
	flag := sliceKit.IsSortedByKey([]string{"a", "aa", "bb", "ccc"}, func(s string) int {
		return len(s)
	})
	fmt.Println(flag) // true
*/
func IsSortedByKey[T any, K constraints.Ordered](s []T, iteratee func(item T) K) bool {
	return lo.IsSortedByKey(s, iteratee)
}
