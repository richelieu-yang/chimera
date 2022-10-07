package setKit

import (
	set "github.com/deckarep/golang-set/v2"
)

// NewSet
/*
PS: Set: 无序；不允许重复.

@param safe	是否goroutines安全？
*/
func NewSet[T comparable](safe bool, args ...T) set.Set[T] {
	if safe {
		return set.NewSet(args...)
	}
	return set.NewThreadUnsafeSet(args...)
}
