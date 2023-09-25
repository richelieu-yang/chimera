package nacosKit

import (
	"testing"
)

func TestMustSetUp(t *testing.T) {
	config := Config{
		NamespaceId: "",
		Addresses:   []string{"http://localhost:8848/nacos"},
	}

	MustSetUp(config)
}
