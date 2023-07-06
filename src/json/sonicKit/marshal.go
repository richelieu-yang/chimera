package sonicKit

import "github.com/bytedance/sonic"

var (
	Marshal func(val interface{}) ([]byte, error) = sonic.Marshal

	MarshalToString func(val interface{}) (string, error) = sonic.MarshalString
)

func MarshalByAPI(api sonic.API, val interface{}) ([]byte, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.Marshal(val)
}

func MarshalToStringByAPI(api sonic.API, val interface{}) (string, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.MarshalToString(val)
}
