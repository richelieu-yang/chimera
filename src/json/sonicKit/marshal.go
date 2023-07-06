package sonicKit

import "github.com/bytedance/sonic"

var (
	Marshal func(val interface{}) ([]byte, error) = sonic.Marshal

	MarshalToString func(val interface{}) (string, error) = sonic.MarshalString
)
