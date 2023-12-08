package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	concurrency := 3
	sem := semaphore.NewWeighted(int64(concurrency))
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if err := worker(id, sem); err != nil {
				logrus.Errorf("Error: %v\n", err)
			}
		}(i)
	}

	wg.Wait()
}

func worker(id int, sem *semaphore.Weighted) error {
	if err := sem.Acquire(context.TODO(), 1); err != nil {
		return err
	}
	defer sem.Release(1)

	// 执行工作任务
	logrus.Infof("Worker %d: Working...\n", id)

	time.Sleep(time.Second)

	return nil
}
