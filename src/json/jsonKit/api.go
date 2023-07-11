package jsonKit

var api API = nil

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}
)

func Marshal(v interface{}) ([]byte, error) {
	return api.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return api.MarshalIndent(v, prefix, indent)
}

func MarshalIndentToString(v interface{}, prefix, indent string) (string, error) {
	data, err := api.MarshalIndent(v, prefix, indent)
	return string(data), err
}

func MarshalToString(v interface{}) (string, error) {
	return api.MarshalToString(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return api.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return api.UnmarshalFromString(str, v)
}
