package sonicKit

import "github.com/bytedance/sonic"

var (
	Unmarshal func(buf []byte, val interface{}) error = sonic.Unmarshal

	UnmarshalString func(buf string, val interface{}) error = sonic.UnmarshalString
)
