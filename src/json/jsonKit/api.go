package jsonKit

var library string

// defaultApi 默认的API
var defaultApi API = nil

// stdApi 标准的API（会对map的keys排序）
var stdApi API = nil

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}
)

func GetLibrary() string {
	return library
}

func GetAPI() API {
	return defaultApi
}
