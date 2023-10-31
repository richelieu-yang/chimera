package mapKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

type (
	MapWithLock[K comparable, V any] struct {
		RWLock mutexKit.RWMutex

		// Map 并发不安全的
		Map map[K]V
	}
)

func NewMapWithLock[K comparable, V any]() *MapWithLock[K, V] {
	return &MapWithLock[K, V]{
		RWLock: mutexKit.RWMutex{},
		Map:    make(map[K]V),
	}
}
