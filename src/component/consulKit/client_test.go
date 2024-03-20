package consulKit

import (
	"github.com/hashicorp/consul/api"
	"testing"
)

func TestNewClient(t *testing.T) {
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500"
	config.Scheme = "http"

	client, err := NewClient(config)
	if err != nil {
		panic(err)
	}
	client = client
}
