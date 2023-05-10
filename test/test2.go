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
	kv := clientv3.NewKV(client)

	key := "/test/"

	// Put
	putOp := clientv3.OpPut(key, "ccc")
	opResp, err := kv.Do(context.TODO(), putOp)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("manager to put")

	// Get
	getOp := clientv3.OpGet(key)
	opResp, err = kv.Do(context.TODO(), getOp)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("manager to get, value: [%s]", string(opResp.Get().Kvs[0].Value))

	// Delete
	delOp := clientv3.OpDelete(key)
	opResp, err = kv.Do(context.TODO(), delOp)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("manager to delete")
}
