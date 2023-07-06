package stackKit

import (
	"container/list"
	"github.com/gogf/gf/v2/container/glist"
)

type (
	stackImpl[V any] struct {
		list *glist.List
	}
)

func (impl *stackImpl[V]) Push(ele V) {
	impl.list.PushBack(ele)
}

func (impl *stackImpl[V]) Pop() (ele V, flag bool) {
	// 可以参考 impl.list.PopBack()，但它只有一个返回值，不满足需求
	// impl.list.PopBack()

	impl.list.LockFunc(func(list *list.List) {
		tmp := list.Back()
		if tmp == nil {
			// 堆栈为空
			return
		}
		ele = list.Remove(tmp).(V)
		flag = true
	})
	return
}

func (impl *stackImpl[V]) Size() int {
	return impl.list.Len()
}
