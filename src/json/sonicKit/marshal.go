package sonicKit

import "github.com/bytedance/sonic"

var (
	Marshal func(v interface{}) ([]byte, error) = sonic.Marshal

	MarshalToString func(v interface{}) (string, error) = sonic.MarshalString
)

func MarshalByAPI(api sonic.API, v interface{}) ([]byte, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.Marshal(v)
}

func MarshalToStringByAPI(api sonic.API, v interface{}) (string, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.MarshalToString(v)
}

func MarshalByAPIWithIndent(api sonic.API, v interface{}, prefix, indent string) ([]byte, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.MarshalIndent(v, prefix, indent)
}

func MarshalToStringByAPIWithIndent(api sonic.API, v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalByAPIWithIndent(api, v, prefix, indent)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
