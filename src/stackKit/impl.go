package stackKit

import (
	"github.com/gogf/gf/v2/container/glist"
)

type (
	stackImpl[V any] struct {
		list *glist.List
	}
)

// Push 放数据（放到slice的最后面）
func (impl *stackImpl[V]) Push(ele V) {
	impl.list.PushBack(ele)
}

// Pop 拿数据（slice最后面的）
func (impl *stackImpl[V]) Pop() V {
	v := impl.list.PopBack()
	return v.(V)
}

func (impl *stackImpl[V]) Size() int {
	return impl.list.Len()
}
