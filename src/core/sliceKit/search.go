package sliceKit

import (
	"github.com/duke-git/lancet/v2/algorithm"
	"github.com/duke-git/lancet/v2/constraints"
)

// BinarySearch （二分法）二分递归查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包constraints.Comparator。
func BinarySearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	return algorithm.BinarySearch(sortedSlice, target, lowIndex, highIndex, comparator)
}

// BinaryIterativeSearch 二分迭代查找，返回元素索引，未找到元素返回-1，参数comparator需要实现包constraints.Comparator。
func BinaryIterativeSearch[T any](sortedSlice []T, target T, lowIndex, highIndex int, comparator constraints.Comparator) int {
	return algorithm.BinaryIterativeSearch(sortedSlice, target, lowIndex, highIndex, comparator)
}

// LinearSearch 基于传入的相等函数线性查找元素，返回元素索引，未找到元素返回-1。
func LinearSearch[T any](slice []T, target T, equal func(a, b T) bool) int {
	return algorithm.LinearSearch(slice, target, equal)
}
