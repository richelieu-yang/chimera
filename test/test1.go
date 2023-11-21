package main

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"github.com/sirupsen/logrus"
	"time"
)

type Interval struct {
	mutexKit.RWMutex

	closed bool

	ticker *time.Ticker

	// closeCh 关闭通道.
	closeCh chan struct{}
}

func (i *Interval) Stop() {
	if i.closed {
		return
	}

	i.LockFunc(func() {
		if i.closed {
			return
		}
		i.closed = true
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
	ticker := time.NewTicker(duration)
	closeCh := make(chan struct{})
	interval := &Interval{
		RWMutex: mutexKit.RWMutex{},
		closed:  false,
		ticker:  ticker,
		closeCh: closeCh,
	}

	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case t := <-ticker.C:
				task(t)
			case <-closeCh:
				return
			}
		}
	}(ticker)
	return interval
}

func main() {
	//创建一个1秒钟间隔的Ticker
	ticker := time.NewTicker(time.Second * 3)

	//启动一个协程来执行代码
	go func() {
		defer func() {
			logrus.Info("ccc")
		}()

		for {
			select {
			case <-ticker.C:
				//每1秒钟执行的代码
				logrus.Info("Ticker running")
			}
		}
	}()

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	ticker.Stop()

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")
}
