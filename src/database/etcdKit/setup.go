package etcdKit

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var kv clientv3.KV
var setupOnce sync.Once

func MustSetUp(config *Config) {
	if err := setUp(config); err != nil {
		logrus.Fatal(err)
	}
}

// setUp
/*
TODO: 可以参考 go-zero 中 registry.go 的 internal.DialClient.

PS:
(1) 如果 Endpoints 无效，会返回error(context.DeadlineExceeded).
*/
func setUp(config *Config) (err error) {
	if err := config.Check(); err != nil {
		return err
	}

	setupOnce.Do(func() {
		var client *clientv3.Client

		v3Config := clientv3.Config{
			Endpoints: config.Endpoints,

			AutoSyncInterval:     time.Minute,
			DialTimeout:          time.Second * 5,
			DialKeepAliveTime:    time.Second * 5,
			DialKeepAliveTimeout: time.Second * 5,
			RejectOldCluster:     true,
			PermitWithoutStream:  true,
		}
		client, err = clientv3.New(v3Config)
		if err == nil {
			// 实例化一个用于操作etcd的KV
			kv = clientv3.NewKV(client)
		}
	})
	return
}

func GetKV() (clientv3.KV, error) {
	if kv == nil {
		return nil, NotSetupError
	}
	return kv, nil
}
