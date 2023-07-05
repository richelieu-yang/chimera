package stackKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"sync"
)

type (
	stackImpl1[V any] struct {
		rwLock *sync.RWMutex

		size int
		eles []V
	}
)

// Push 放数据（放到slice的最后面）
func (s *stackImpl1[V]) Push(ele V) {
	if s.rwLock != nil {
		s.rwLock.Lock()
		defer s.rwLock.Unlock()
	}

	s.eles = append(s.eles, ele)
	s.size++
}

// Pop 拿数据（slice最后面的）
func (s *stackImpl1[V]) Pop() (ele V, ok bool) {
	if s.rwLock != nil {
		s.rwLock.Lock()
		defer s.rwLock.Unlock()
	}

	s.eles, ele, ok = sliceKit.RemoveLast(s.eles)
	if ok {
		s.size--
	}
	return
}

func (s *stackImpl1[V]) Size() int {
	if s.rwLock != nil {
		s.rwLock.RLock()
		defer s.rwLock.RUnlock()
	}

	return s.size
}
