package stackKit

import "sync"

// NewStack 堆栈
/*
@param safe	是否goroutines安全？
@return 	必定不为nil
*/
func NewStack[V any](safe bool) Stack[V] {
	stack := &stackImpl[V]{
		size: 0,
		eles: make([]V, 0, 32),
	}
	if safe {
		stack.rwLock = new(sync.RWMutex)
	}
	return stack
}
