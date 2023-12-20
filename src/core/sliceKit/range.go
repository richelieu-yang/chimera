package sliceKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"golang.org/x/exp/constraints"
)

// Range 根据指定的起始值和数量，创建一个数字切片（step为1）。
/*
e.g.
	fmt.Println(sliceKit.Range(-4, 4)) // [-4 -3 -2 -1]
*/
func Range[T constraints.Integer | constraints.Float](start T, count int) []T {
	return mathutil.Range(start, count)
}

// RangeWithStep 根据指定的起始值，结束值，步长，创建一个数字切片。
/*
e.g.
	fmt.Println(sliceKit.RangeWithStep(-4, 1, 2)) // [-4 -2 0]
*/
func RangeWithStep[T constraints.Integer | constraints.Float](start, end, step T) []T {
	return mathutil.RangeWithStep(start, end, step)
}
