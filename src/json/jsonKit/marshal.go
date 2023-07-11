package jsonKit

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

// MarshalByAPIWithIndent
/*
@param prefix 一般为""
@param indent 推荐值: "\t" || "    "（4个空格）
*/
func MarshalByAPIWithIndent(api sonic.API, v interface{}, prefix, indent string) ([]byte, error) {
	if api == nil {
		api = sonic.ConfigDefault
	}
	return api.MarshalIndent(v, prefix, indent)
}

// MarshalToStringByAPIWithIndent
/*
@param prefix 一般为""
@param indent 推荐值: "\t" || "    "（4个空格）
*/
func MarshalToStringByAPIWithIndent(api sonic.API, v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalByAPIWithIndent(api, v, prefix, indent)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
