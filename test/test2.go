package main

import (
	"context"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	//kv := clientv3.NewKV(client)

	leaseGrantResponse, err := client.Grant(context.TODO(), 10)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("lease id: [%d].", leaseGrantResponse.ID)

	ch, err := client.KeepAlive(context.TODO(), leaseGrantResponse.ID)
	if err != nil {
		logrus.Fatal(err)
	}
	go func() {
		defer func() {
			logrus.Info("goroutine ends")
		}()

		for {
			select {
			case keepResp := <-ch:
				if keepResp == nil {
					// e.g. 该租约被删除了
					logrus.Warn("租约已失效")
					goto END
				}
				logrus.WithFields(logrus.Fields{
					"ID":  keepResp.ID,
					"TTL": keepResp.TTL,
				}).Info("续租成功.")
			}
		}
	END:
	}()

	select {}
}
