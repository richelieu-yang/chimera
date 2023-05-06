package main

import (
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
		logrus.Panic(err)
	}

}
