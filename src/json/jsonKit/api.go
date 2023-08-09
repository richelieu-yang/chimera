package jsonKit

var defaultAPI API = nil

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}
)

func GetAPI() API {
	return defaultAPI
}

func SetAPI(api API) {
	defaultAPI = api
}

var (
	Marshal func(v interface{}) ([]byte, error) = defaultAPI.Marshal

	MarshalIndent func(v interface{}, prefix, indent string) ([]byte, error) = defaultAPI.MarshalIndent

	MarshalToString func(v interface{}) (string, error) = defaultAPI.MarshalToString

	Unmarshal func(data []byte, v interface{}) error = defaultAPI.Unmarshal

	UnmarshalFromString func(str string, v interface{}) error = defaultAPI.UnmarshalFromString
)

// MarshalIndentToString
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentToString(v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalIndent(v, prefix, indent)
	return string(data), err
}
