package sonicKit

import "github.com/bytedance/sonic"

var (
	Marshal func(val interface{}) ([]byte, error) = sonic.Marshal

	MarshalString func(val interface{}) (string, error) = sonic.MarshalString
)
