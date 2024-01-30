package mathKit

import "math"

var (
	Inf func(sign int) float64 = math.Inf

	IsInf func(f float64, sign int) bool = math.IsInf
)
