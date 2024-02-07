package setKit

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
)

type (
	SetWithLock[T comparable] struct {
		mutexKit.RWMutex

		// Set 并发不安全的
		Set mapset.Set[T]
	}
)

func (set *SetWithLock[E]) Size() (size int) {
	if set == nil {
		return
	}

	/* 读锁 */
	set.RLockFunc(func() {
		size = set.Set.Cardinality()
	})
	return
}

func NewSetWithLock[T comparable]() *SetWithLock[T] {
	return &SetWithLock[T]{
		Set: NewSet[T](false), // 并发不安全的
	}
}
