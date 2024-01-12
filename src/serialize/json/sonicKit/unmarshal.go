package sonicKit

import "github.com/bytedance/sonic"

var (
	Unmarshal func(buf []byte, val interface{}) error = sonic.Unmarshal

	UnmarshalFromString func(buf string, val interface{}) error = sonic.UnmarshalString
)

func UnmarshalByAPI(api sonic.API, data []byte, v interface{}) error {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.Unmarshal(data, v)
}

func UnmarshalFromStringByAPI(api sonic.API, str string, v interface{}) error {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.UnmarshalFromString(str, v)
}
