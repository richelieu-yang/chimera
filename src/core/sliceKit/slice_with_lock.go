package sliceKit

import (
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

type (
	SliceWithLock[E any] struct {
		mutexKit.RWMutex

		// Slice 并发不安全的
		Slice []E
	}
)

func (s *SliceWithLock[E]) Size() (size int) {
	if s == nil {
		return
	}

	/* 读锁 */
	s.RLockFunc(func() {
		size = len(s.Slice)
	})
	return
}

func NewSliceWithLock[E any]() *SliceWithLock[E] {
	return &SliceWithLock[E]{
		Slice: make([]E, 0, 8),
	}
}
