/*
Algorithm
algorithm 算法包实现一些基本算法，sort、lrucache.

https://www.golancet.cn/api/packages/algorithm.html#Algorithm
*/

package sliceKit

import (
	"github.com/duke-git/lancet/v2/algorithm"
	"github.com/duke-git/lancet/v2/constraints"
)

// BubbleSort 冒泡排序，参数comparator需要实现包constraints.Comparator。
func BubbleSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.BubbleSort(slice, comparator)
}

// InsertionSort 插入排序，参数comparator需要实现包constraints.Comparator。
func InsertionSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.InsertionSort(slice, comparator)
}

// SelectionSort 选择排序，参数comparator需要实现包constraints.Comparator。
func SelectionSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.SelectionSort(slice, comparator)
}

// ShellSort 希尔排序，参数comparator需要实现包constraints.Comparator。
func ShellSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.ShellSort(slice, comparator)
}

// QuickSort 快速排序，参数comparator需要实现包constraints.Comparator。
func QuickSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.QuickSort(slice, comparator)
}

// HeapSort 堆排序，参数comparator需要实现包constraints.Comparator。
func HeapSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.HeapSort(slice, comparator)
}

// MergeSort 归并排序，参数comparator需要实现包constraints.Comparator。
func MergeSort[T any](slice []T, comparator constraints.Comparator) {
	algorithm.MergeSort(slice, comparator)
}

// CountSort 计数排序，参数comparator需要实现包constraints.Comparator。
func CountSort[T any](slice []T, comparator constraints.Comparator) []T {
	return algorithm.CountSort(slice, comparator)
}
