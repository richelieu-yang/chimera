package redisKit

import "testing"

func TestSetUp(t *testing.T) {
	config := &Config{
		UserName: "",
		Password: "",
		Mode:     ModeSingleNode,
		SingleNodeConfig: &SingleNodeConfig{
			Addr: "127.0.0.1:6379",
			DB:   0,
		},
	}
	MustSetUp(config)
}
