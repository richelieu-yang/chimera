package mathKit

import "golang.org/x/exp/constraints"

// Max
/*
TODO: go1.21，使用max().
*/
func Max[T constraints.Ordered](p T, args ...T) T {
	rst := p

	for _, ele := range args {
		if ele > rst {
			rst = ele
		}
	}
	return rst
}

// Min
/*
TODO: go1.21，使用min().
*/
func Min[T constraints.Ordered](p T, args ...T) T {
	rst := p

	for _, ele := range args {
		if ele < rst {
			rst = ele
		}
	}
	return rst
}
