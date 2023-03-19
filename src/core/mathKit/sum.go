package mathKit

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Sum 求和
/*

 */
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](s []T) T {
	return lo.Sum(s)
}

// SumBy 求和
/*

 */
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](s []T, iteratee func(item T) R) R {
	return lo.SumBy(s, iteratee)
}
