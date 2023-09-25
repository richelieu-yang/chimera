package nacosKit

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8849/nacos111"},
	}

	MustSetUp(config)

	cc, err := NewConfigClient()
	if err != nil {
		panic(err)
	}
	defer cc.CloseClient()

	text, err := cc.GetConfig(vo.ConfigParam{
		DataId: "2",
		//Group:  "group",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(text)

	//cc.ListenConfig(vo.ConfigParam{
	//	DataId: "",
	//	Group:  "",
	//	OnChange: func(namespace, group, dataId, data string) {
	//
	//	},
	//})
}
