package nacosKit

import (
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := Config{
		NamespaceId: "",
		//Addresses:   []string{"http://localhost:8848/nacos"},
		Addresses: []string{"http://192.168.60.206:8849/nacos"},
	}

	MustSetUp(config)
}
