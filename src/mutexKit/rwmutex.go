package mutexKit

import "sync"

type (
	RWMutex struct {
		sync.RWMutex
	}
)

func NewRWMutex() *RWMutex {
	return &RWMutex{sync.RWMutex{}}
}

// LockFunc 写锁
func (m *RWMutex) LockFunc(f func()) {
	m.Lock()
	defer m.Unlock()

	f()
}

// RLockFunc 读锁
func (m *RWMutex) RLockFunc(f func()) {
	m.RLock()
	defer m.RUnlock()

	f()
}

// TryLockFunc 写锁
/*
PS: 不管加锁成功还是失败，不会阻塞.
*/
func (m *RWMutex) TryLockFunc(f func()) (result bool) {
	if m.TryLock() {
		defer m.Unlock()

		result = true
		f()
	}
	return
}

// TryRLockFunc 读锁
/*
PS: 不管加锁成功还是失败，不会阻塞.
*/
func (m *RWMutex) TryRLockFunc(f func()) (result bool) {
	if m.TryRLock() {
		defer m.RUnlock()

		result = true
		f()
	}
	return
}
