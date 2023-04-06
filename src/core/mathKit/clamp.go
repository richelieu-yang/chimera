package mathKit

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Clamp
/*
value < min: 	返回min
value > max: 	返回max
others:			返回value

e.g.
(0, -10, 10)	=> 0
(-42, -10, 10)	=> -10
(42, -10, 10)	=> 10
*/
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	return lo.Clamp(value, min, max)
}
