package nacosKit

import (
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8848/nacos"},
		//Addresses:   []string{"http://localhost:8849/nacos111"},
	}

	MustSetUp(config)

	cc, err := NewConfigClient()
	if err != nil {
		panic(err)
	}
	defer cc.CloseClient()

	//nc, err := NewNamingClient()
	//if err != nil {
	//	panic(err)
	//}
	//defer nc.CloseClient()

	//ulid := idKit.NewULID()
	//ulid = ulid
	//text, err := cc.GetConfig(vo.ConfigParam{
	//	//DataId: "1",
	//	DataId: ulid,
	//	//Group:  ulid,
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(text)

	//cc.ListenConfig(vo.ConfigParam{
	//	DataId: "",
	//	Group:  "",
	//	OnChange: func(namespace, group, dataId, data string) {
	//
	//	},
	//})
}
