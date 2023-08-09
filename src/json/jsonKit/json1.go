package jsonKit

func MarshalWithAPI(api API, v interface{}) ([]byte, error) {
	if api == nil {
		api = defaultAPI
	}

	return api.Marshal(v)
}

// MarshalIndentWithAPI
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentWithAPI(api API, v interface{}, prefix, indent string) ([]byte, error) {
	if api == nil {
		api = defaultAPI
	}

	return api.MarshalIndent(v, prefix, indent)
}

func MarshalToStringWithAPI(api API, v interface{}) (string, error) {
	if api == nil {
		api = defaultAPI
	}

	return api.MarshalToString(v)
}

// MarshalIndentToStringWithAPI
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentToStringWithAPI(api API, v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalIndentWithAPI(api, v, prefix, indent)
	return string(data), err
}

func UnmarshalWithAPI(api API, data []byte, v interface{}) error {
	if api == nil {
		api = defaultAPI
	}

	return api.Unmarshal(data, v)
}

func UnmarshalFromStringWithAPI(api API, str string, v interface{}) error {
	if api == nil {
		api = defaultAPI
	}

	return api.UnmarshalFromString(str, v)
}
