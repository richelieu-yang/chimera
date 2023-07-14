// go版本 < 1.21
//go:build !go1.21

package mathKit

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](p T, args ...T) T {
	rst := p

	for _, ele := range args {
		if ele > rst {
			rst = ele
		}
	}
	return rst
}

func Min[T constraints.Ordered](p T, args ...T) T {
	rst := p

	for _, ele := range args {
		if ele < rst {
			rst = ele
		}
	}
	return rst
}
