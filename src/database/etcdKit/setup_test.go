package etcdKit

import (
	"context"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/core/osKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	var path string
	if osKit.IsWindows() {
		path = "D:\\GolandProjects\\chimera\\chimera-lib\\config.yaml"
	} else {
		path = "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"
	}

	type config struct {
		Etcd *Config `json:"etcd"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	MustSetUp(c.Etcd)

	kv, err := GetKV()
	if err != nil {
		logrus.Fatal(err)
	}
	resp, err := kv.Put(context.TODO(), "k", "v")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(resp.Header.GetRevision())
}
