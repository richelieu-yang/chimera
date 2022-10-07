package sliceKit

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Sort （简单数据类型）排序
/*
@param s 	可以为nil，将返回nil
@param args true: 降序；false: 升序（默认）
@return 	可能为nil

e.g.
[int](nil)    	 				=> nil
[int](nil, true)				=> nil
([]int{0, 1, 9, 8, 1}) 		 	=> [0 1 1 8 9]
([]int{0, 1, 9, 8, 1}, true) 	=> [9 8 1 1 0]
*/
func Sort[T constraints.Ordered](s []T, args ...bool) []T {
	var less func(i, j int) bool
	byDescending := GetFirstItemWithDefault(false, args...)
	if byDescending {
		// 降序
		less = func(i, j int) bool {
			return s[i] > s[j]
		}
	} else {
		// 升序
		less = func(i, j int) bool {
			return s[i] < s[j]
		}
	}

	sort.Slice(s, less)
	return s
}

// SortComplexSlice 排序复杂类型的slice实例（元素的类型可以为自定义结构体）
/*
PS:
(1) 需要注意的是 sort.Sort() 函数不能保证数据排序是稳定的，
	如果需要保证数据排序稳定，可以使用 sort.Stable() 函数。（“稳定”的含义是原始数据中 a 和 b 的值相等，排序前 a 排在 b 的前面，排序后 a 仍排在 b 的前面）
(2) 现成的sort.Interface接口的实现: sort.IntSlice、sort.Float64Slice、sort.StringSlice.

@param data	可以为nil，将返回nil
@param args true: 稳定; false: 不稳定（默认；大部分场景这就够了）
@return 	可能为nil

e.g.
s := []int{0, 2, 3, 1}
sliceKit.SortComplexSlice(sort.IntSlice(s))
fmt.Println(s) 	// [0 1 2 3]
*/
func SortComplexSlice(data sort.Interface, args ...bool) sort.Interface {
	if data == nil {
		return nil
	}

	stable := GetFirstItemWithDefault(false, args...)
	if stable {
		sort.Stable(data)
	} else {
		sort.Sort(data)
	}
	// 此时返回值必定不为nil
	return data
}
