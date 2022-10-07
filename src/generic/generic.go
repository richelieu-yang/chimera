package generic

import "golang.org/x/exp/constraints"

// Number 所有整数与浮点数
type Number interface {
	constraints.Integer | constraints.Float
}
