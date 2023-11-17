package gaodeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

type (
	Client struct {
		key string
	}
)

func NewClient(key string) (*Client, error) {
	if err := strKit.AssertNotBlank(key, "key"); err != nil {
		return nil, err
	}
	return &Client{
		key: key,
	}, nil
}
