package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := &Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8848/nacos"},
	}

	MustSetUp(config)

	namingClient, err := NewNamingClient()
	if err != nil {
		panic(err)
	}
	defer namingClient.CloseClient()
	configClient, err := NewConfigClient()
	if err != nil {
		panic(err)
	}
	defer configClient.CloseClient()

	configClient.GetConfig(vo.ConfigParam{
		DataId: "common",
		//Group:  "",
	})

}
