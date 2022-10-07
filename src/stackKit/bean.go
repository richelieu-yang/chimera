package stackKit

import (
	"gitee.com/richelieu042/go-scales/src/core/sliceKit"
	"sync"
)

type (
	Stack[V any] interface {
		Push(ele V)

		// Pop
		/*
			@return 被拿出来的元素 + 是否成功拿出数据？
		*/
		Pop() (V, bool)

		Size() int
	}

	// stackImpl 堆栈（先进后出）
	stackImpl[V any] struct {
		rwLock *sync.RWMutex
		size   int
		eles   []V
	}
)

// Push 放数据（放到slice的最后面）
func (s *stackImpl[V]) Push(ele V) {
	if s.rwLock != nil {
		s.rwLock.Lock()
		defer s.rwLock.Unlock()
	}

	s.eles = append(s.eles, ele)
	s.size++
}

// Pop 拿数据（slice最后面的）
func (s *stackImpl[V]) Pop() (ele V, ok bool) {
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

func (s *stackImpl[V]) Size() int {
	if s.rwLock != nil {
		s.rwLock.RLock()
		defer s.rwLock.RUnlock()
	}

	return s.size
}
