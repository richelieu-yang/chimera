package main

import (
	"context"
	"fmt"
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
	// 实例化一个用于操作etcd的KV
	kv := clientv3.NewKV(client)

	getResp, err := kv.Get(context.TODO(), "/school/class/students")
	if err != nil {
		logrus.Panic(err)
	}
	// 输出本次的Revision
	fmt.Printf("Key is s %s \n Value is %s \n", getResp.Kvs[0].Key, getResp.Kvs[0].Value)

}
