package mutexKit

import (
	"sync"
)

// Mutex 互斥锁.
/*
PS: 结构体中使用此锁，可以考虑"匿名字段".
*/
type Mutex struct {
	sync.Mutex
}

func (m *Mutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()

	f()
}

func (m *Mutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		defer m.Unlock()

		result = true
		f()
	}
	return
}

// NewMutex Deprecated: 直接用 &Mutex{} 吧.
func NewMutex() *Mutex {
	return &Mutex{}
}
