package sliceKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

type (
	SliceWithLock[E any] struct {
		RWLock mutexKit.RWMutex

		// Slice 并发不安全的
		Slice []E
	}
)

func NewSliceWithLock[E any]() *SliceWithLock[E] {
	return &SliceWithLock[E]{
		RWLock: mutexKit.RWMutex{},
		Slice:  make([]E, 0, 8),
	}
}
