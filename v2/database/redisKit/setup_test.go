package redisKit

import (
	"context"
	"github.com/richelieu42/chimera/v2/core/timeKit"
	"github.com/richelieu42/chimera/v2/idKit"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	key := idKit.NewULID()
	value := timeKit.FormatCurrentTime()
	_, err := client.Set(context.TODO(), key, value, 0)
	assert.Nil(t, err)
	_, err = client.Del(context.TODO(), key)
	assert.Nil(t, err)
}
