package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"sync"
)

func PushToAll(data []byte, exceptBsids []string) (err error) {
	if err = isAvailable(); err != nil {
		return err
	}

	/* map读锁 */
	idMap.RLockFunc(func() {
		var wg sync.WaitGroup
		for _, channel := range idMap.Map {
			c := channel

			if sliceKit.Contains(exceptBsids, c.GetBsid()) {
				continue
			}
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

func PushToBsid(data []byte, bsid string) (err error) {
	if err = isAvailable(); err != nil {
		return
	}

	/* map读锁 */
	bsidMap.RLockFunc(func() {
		channel := bsidMap.Map[bsid]
		if channel == nil {
			err = errorKit.New("No channel for bsid(%s)", bsid)
			return
		}
		err = channel.Push(data)
	})
	return
}

func PushToUser(data []byte, user string, exceptBsids []string) (err error) {
	if err = isAvailable(); err != nil {
		return
	}

	/* map读锁 */
	userMap.RLockFunc(func() {
		userSet := userMap.Map[user]
		if userSet == nil {
			err = errorKit.New("No channels for user(%s)", user)
			return
		}

		/* set读锁 */
		userSet.RLockFunc(func() {
			var wg sync.WaitGroup
			userSet.Set.Each(func(channel Channel) bool {
				c := channel

				if sliceKit.Contains(exceptBsids, c.GetBsid()) {
					return false // 不中断循环
				}
				wg.Add(1)
				_ = pool.Submit(func() {
					defer wg.Done()
					_ = c.Push(data)
				})
				return false // 不中断循环
			})
			wg.Wait()
		})
	})
	return
}

func PushToGroup(data []byte, group string, exceptBsids []string) (err error) {
	if err = isAvailable(); err != nil {
		return
	}

	/* map读锁 */
	groupMap.RLockFunc(func() {
		groupSet := userMap.Map[group]
		if groupSet == nil {
			err = errorKit.New("No channels for group(%s)", group)
			return
		}

		/* set读锁 */
		groupSet.RLockFunc(func() {
			var wg sync.WaitGroup
			groupSet.Set.Each(func(channel Channel) bool {
				c := channel

				if sliceKit.Contains(exceptBsids, c.GetBsid()) {
					return false // 不中断循环
				}
				wg.Add(1)
				_ = pool.Submit(func() {
					defer wg.Done()
					_ = c.Push(data)
				})
				return false // 不中断循环
			})
			wg.Wait()
		})
	})
	return
}
