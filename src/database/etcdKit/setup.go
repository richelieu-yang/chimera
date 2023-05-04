package etcdKit

import (
	"context"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var client *clientv3.Client
var setupOnce sync.Once

func MustSetUp(config *Config) {
	if err := setUp(config); err != nil {
		logrus.Info(err == context.DeadlineExceeded)
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
	if config == nil {
		err = errorKit.Simple("config == nil")
		return
	}

	setupOnce.Do(func() {
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
	})
	return
}

func GetClient() (*clientv3.Client, error) {
	if client == nil {
		return nil, NotSetupError
	}
	return client, nil
}
