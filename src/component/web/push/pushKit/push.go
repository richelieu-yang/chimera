package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"sync"
)

func PushToAll(data []byte, exceptBsids []string) (err error) {
	if err = isAvailable(); err != nil {
		return err
	}

	/* 写锁 */
	idMap.RWLock.LockFunc(func() {
		var wg sync.WaitGroup

		for _, channel := range idMap.Map {
			if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
				continue
			}

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

func PushToUser(data []byte, user string, exceptBsids []string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	userSet := GetUserSet(user)
	if userSet == nil {
		return errorKit.New("No set for user(%s)", user)
	}

	/* 写锁 */
	userSet.RWLock.LockFunc(func() {
		var wg sync.WaitGroup
		userSet.Set.Each(func(channel types.Channel) bool {
			if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
				return false // 不中断循环
			}

			c := channel
			wg.Add(1)
			_ = pool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
			return false // 不中断循环
		})
	})
	return nil
}

func PushToGroup(data []byte, group string, exceptBsids []string) error {
	if err := isAvailable(); err != nil {
		return err
	}

	groupSet := GetGroupSet(group)
	if groupSet == nil {
		return errorKit.New("No set for group(%s)", group)
	}

	/* 写锁 */
	groupSet.RWLock.LockFunc(func() {
		var wg sync.WaitGroup
		groupSet.Set.Each(func(channel types.Channel) bool {
			if sliceKit.Contains(exceptBsids, channel.GetBsid()) {
				return false // 不中断循环
			}

			c := channel
			wg.Add(1)
			_ = pool.Submit(func() {
				defer wg.Done()
				_ = c.Push(data)
			})
			return false // 不中断循环
		})
	})
	return nil
}
