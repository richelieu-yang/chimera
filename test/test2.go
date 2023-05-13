package main

import (
	"context"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"sync"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)

	v3Config := clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},

		AutoSyncInterval:     time.Minute,
		DialTimeout:          time.Second * 5,
		DialKeepAliveTime:    time.Second * 5,
		DialKeepAliveTimeout: time.Second * 5,
		RejectOldCluster:     true,
		PermitWithoutStream:  true,
	}
	client, err := clientv3.New(v3Config)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Close()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		session, err := concurrency.NewSession(client)
		if err != nil {
			logrus.Fatal(err)
		}
		defer session.Close()
		mu := concurrency.NewMutex(session, "/lock")

		// 加锁
		logrus.Info("[goroutine 0] ready to lock on")
		if err := mu.Lock(context.TODO()); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("[goroutine 0] lock on")

		// 等一会
		time.Sleep(time.Second * 3)

		// 解锁
		logrus.Info("[goroutine 0] ready to lock off")
		if err := mu.Unlock(context.TODO()); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("[goroutine 0] lock off")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		session, err := concurrency.NewSession(client)
		if err != nil {
			logrus.Fatal(err)
		}
		defer session.Close()
		mu := concurrency.NewMutex(session, "/lock")

		// 等一会，让锁被抢掉
		time.Sleep(time.Second)

		// 加锁
		logrus.Info("[goroutine 1] ready to lock on")
		if err := mu.Lock(context.TODO()); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("[goroutine 1] lock on")

		// 等一会
		time.Sleep(time.Second * 3)

		// 解锁
		logrus.Info("[goroutine 1] ready to lock off")
		if err := mu.Unlock(context.TODO()); err != nil {
			logrus.Fatal(err)
		}
		logrus.Info("[goroutine 1] lock off")
	}()

	wg.Wait()
	logrus.Info("======")
}
