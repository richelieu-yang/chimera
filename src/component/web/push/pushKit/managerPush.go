package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"sync"
)

func PushToAll(data []byte) (err error) {
	if err = isAvailable(); err != nil {
		return err
	}

	// 写锁
	allMap.RWLock.LockFunc(func() {
		var wg sync.WaitGroup

		for _, channel := range allMap.Map {
			c := channel
			wg.Add(1)
			_ = pool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
		}
		wg.Wait()
	})

	return nil
}

func PushToGroup(data []byte, group string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	groupSet := GetGroupSet(group)
	if groupSet == nil {
		return errorKit.New("No set for group(%s)", group)
	}

	// 写锁
	groupSet.RWLock.LockFunc(func() {
		var wg sync.WaitGroup
		groupSet.Set.Each(func(channel Channel) bool {
			c := channel
			wg.Add(1)
			_ = pool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
			// 不中断循环
			return false
		})
	})
	return nil
}

func PushToUser(data []byte, user string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	userSet := GetUserSet(user)
	if userSet == nil {
		return errorKit.New("No set for user(%s)", user)
	}

	// 写锁
	userSet.RWLock.LockFunc(func() {
		var wg sync.WaitGroup
		userSet.Set.Each(func(channel Channel) bool {
			c := channel
			wg.Add(1)
			_ = pool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
			// 不中断循环
			return false
		})
	})
	return nil
}

func PushToBsid(data []byte, bsid string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	channel := GetChannelByBsid(bsid)
	if channel == nil {
		return errorKit.New("No channel for bsid(bsid)", bsid)
	}
	return channel.Push(data)
}
