package setKit

import (
	set "github.com/deckarep/golang-set/v2"
)

// NewSet
/*
PS:
(1) Set: 无序；不允许重复.

@param threadSafe	是否goroutines安全？
*/
func NewSet[T comparable](threadSafe bool, args ...T) set.Set[T] {
	if threadSafe {
		return set.NewSet(args...)
	}
	return set.NewThreadUnsafeSet(args...)
}

func NewSetFromMapKeys[T comparable, V any](threadSafe bool, val map[T]V) set.Set[T] {
	if threadSafe {
		return set.NewSetFromMapKeys(val)
	}
	return set.NewThreadUnsafeSetFromMapKeys(val)
}
