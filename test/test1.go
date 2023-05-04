package main

import (
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)

	// "127.0.0.1:12379" 地址是无效的
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},

		// 参考go-zero中的 internal.DialClient
		AutoSyncInterval:     time.Minute,
		DialTimeout:          time.Second * 5,
		DialKeepAliveTime:    time.Second * 5,
		DialKeepAliveTimeout: time.Second * 5,
		RejectOldCluster:     true,
		PermitWithoutStream:  true,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Close()

	//ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	//defer cancel()
	//_, err = client.Get(ctx, "key")
	//if err != nil {
	//	logrus.Infof("err == context.DeadlineExceeded? [%t]", err == context.DeadlineExceeded)
	//	logrus.Fatal(err)
	//}
}
