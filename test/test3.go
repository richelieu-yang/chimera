package main

import (
	"github.com/richelieu-yang/chimera/v2/src/component/nacosKit"
)

func main() {
	config := nacosKit.Config{
		NamespaceId: "",
		//Addresses:   []string{"http://localhost:8848/nacos"},
		Addresses: nil,
	}

	nacosKit.MustSetUp(config)
}
