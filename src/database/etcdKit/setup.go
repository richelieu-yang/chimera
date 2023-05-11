package etcdKit

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/log/zapKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

var client *clientv3.Client
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
	if err = config.Check(); err != nil {
		return
	}

	setupOnce.Do(func() {
		/* etcd客户端日志输出 */
		var logger *zap.Logger
		if strKit.IsNotEmpty(config.LogPath) {
			var writer io.Writer
			writer, err = ioKit.NewLumberjackWriteCloser(ioKit.WithFilePath(config.LogPath))
			if err != nil {
				return
			}
			logger, err = zapKit.NewLogger(writer, zap.InfoLevel)
			if err != nil {
				return
			}
		}

		v3Config := clientv3.Config{
			Endpoints: config.Endpoints,
			Logger:    logger,

			AutoSyncInterval:     time.Minute,
			DialTimeout:          time.Second * 5,
			DialKeepAliveTime:    time.Second * 5,
			DialKeepAliveTimeout: time.Second * 5,
			RejectOldCluster:     true,
			PermitWithoutStream:  true,
		}
		client, err = clientv3.New(v3Config)
		if err != nil {
			client = nil
		}
	})
	return
}

// GetClient
/*
PS:
(1) 要使用 KV 的情况下，建议调用 clientv3.NewKV() 以实例化一个用于操作etcd的KV（内置错误重试机制）.
(2) 租约相关需要用到 *clientv3.Client实例.
*/
func GetClient() (*clientv3.Client, error) {
	if client == nil {
		return nil, NotSetupError
	}
	return client, nil
}
