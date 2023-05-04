package etcdKit

import "testing"

func TestMustSetUp(t *testing.T) {
	MustSetUp(&Config{Endpoints: []string{
		"127.0.0.1:12379",
	}})
}
