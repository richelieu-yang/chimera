package mutexKit

import (
	"sync"
)

type (
	Mutex struct {
		*sync.Mutex
	}
)

func NewMutex() *Mutex {
	return &Mutex{&sync.Mutex{}}
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
