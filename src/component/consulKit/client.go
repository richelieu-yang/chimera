package consulKit

import (
	"github.com/hashicorp/consul/api"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
)

// NewClient
/*
@param config 	(1) 不能为 nil;
				(2) Address 默认为: "127.0.0.1:8500"，Scheme 默认为: "http";
				(3) 一般情况下，配置下 Address 即可.
*/
func NewClient(config *api.Config) (*api.Client, error) {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return nil, err
	}

	return api.NewClient(config)
}
