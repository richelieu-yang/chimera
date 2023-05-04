package etcdKit

import (
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/core/osKit"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	var path string
	if osKit.IsWindows() {
		path = ""
	} else {
		path = "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"
	}

	type config struct {
		Etcd *Config `json:"etcd"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	MustSetUp(c.Etcd)

	client, err := GetClient()
	client.Put()
}
