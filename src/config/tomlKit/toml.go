package tomlKit

import "github.com/pelletier/go-toml/v2"

var (
	Unmarshal func(data []byte, v interface{}) error = toml.Unmarshal
)
