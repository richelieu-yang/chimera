package timeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"time"
)

type Interval struct {
	mutexKit.RWMutex

	stopped bool

	ticker *time.Ticker

	// closeCh 关闭通道.
	closeCh chan struct{}
}

// Stop
/*
PS:
(1) 可以多次调用，不会panic，但这样没意义，调用一次就够了;
(2) 如果有任务正在执行，会等它先执行完.
*/
func (i *Interval) Stop() {
	if i == nil || i.stopped {
		return
	}

	/* 写锁 */
	i.LockFunc(func() {
		if i.stopped {
			return
		}
		i.stopped = true
		i.ticker.Stop()
		i.closeCh <- struct{}{}
	})
}

// NewInterval
/*
@param task		不能为nil
@param duration 必须>0
*/
func NewInterval(task func(t time.Time), duration time.Duration) *Interval {
	i := &Interval{
		RWMutex: mutexKit.RWMutex{},
		stopped: false,
		ticker:  time.NewTicker(duration),
		closeCh: make(chan struct{}),
	}

	go func(i *Interval) {
		//// test
		//defer func() {
		//	logrus.Info("goroutine ends")
		//}()

		defer i.ticker.Stop()

		for {
			select {
			case t := <-i.ticker.C:
				/* 读锁 */
				i.RLockFunc(func() {
					if i.stopped {
						return
					}
					task(t)
				})
			case <-i.closeCh:
				return
			}
		}
	}(i)
	return i
}

var SetInterval func(task func(t time.Time), duration time.Duration) *Interval = NewInterval

// ClearInterval
/*
@param i 可以为nil
*/
func ClearInterval(i *Interval) {
	i.Stop()
}
