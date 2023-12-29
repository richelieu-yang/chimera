package compareKit

import (
	"github.com/duke-git/lancet/v2/compare"
	"golang.org/x/exp/constraints"
)

var (
	// Equal 检查两个值是否相等(检查类型和值).
	Equal func(left, right any) bool = compare.Equal

	// EqualValue 检查两个值是否相等(只检查值).
	EqualValue func(left, right any) bool = compare.EqualValue

	// LessThan 验证参数`left`的值是否小于参数`right`的值.
	LessThan func(left, right any) bool = compare.LessThan

	// GreaterThan 验证参数`left`的值是否大于参数`right`的值.
	GreaterThan func(left, right any) bool = compare.GreaterThan

	// LessOrEqual 验证参数`left`的值是否小于或等于参数`right`的值.
	LessOrEqual func(left, right any) bool = compare.LessOrEqual

	// GreaterOrEqual 验证参数`left`的值是否大于或参数`right`的值.
	GreaterOrEqual func(left, right any) bool = compare.GreaterOrEqual
)

// InDelta 检查增量内两个值是否相等.
func InDelta[T constraints.Integer | constraints.Float](left, right T, delta float64) bool {
	return compare.InDelta(left, right, delta)
}
