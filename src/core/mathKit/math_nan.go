package mathKit

import "math"

var (
	NaN func() float64 = math.NaN

	IsNaN func(f float64) (is bool) = math.IsNaN
)
