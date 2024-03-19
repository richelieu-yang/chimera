package consulKit

import (
	"github.com/hashicorp/consul/api"
)

func NewClient() (*api.Client, error) {
	cfg := api.DefaultConfig()
	cfg.Address = ""

	return api.NewClient(cfg)
}
