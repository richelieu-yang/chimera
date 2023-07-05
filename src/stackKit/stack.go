package stackKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/listKit"
)

type (
	Stack[V any] interface {
		// Push 放在最后面
		Push(ele V)

		// Pop 移除并返回最后面的
		Pop() V

		Size() int
	}
)

func NewStack[V any](safe ...bool) Stack[V] {
	return &stackImpl[V]{
		list: listKit.NewDoubleLinkedList(safe...),
	}
}

// NewStackFrom
/*
Deprecated: 有点耗性能，看 GoFrame 后续会不会支持泛型.
*/
func NewStackFrom[V any](s []V, safe ...bool) Stack[V] {
	var s1 []interface{} = sliceKit.ConvertElementType[V, interface{}](s, func(item V, index int) interface{} {
		return item
	})
	return &stackImpl[V]{
		list: listKit.NewDoubleLinkedListFrom(s1, safe...),
	}
}

//// NewStack1 堆栈（后进先出）
///*
//@param safe	是否goroutines安全？
//@return 必定不为nil
//*/
//func NewStack1[V any](safe bool) Stack[V] {
//	stack := &stackImpl1[V]{
//		size: 0,
//		eles: make([]V, 0, 32),
//	}
//	if safe {
//		stack.rwLock = new(sync.RWMutex)
//	}
//	return stack
//}
