package jsonResplKit

type (
	API interface {
		Marshal(v interface{}) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		Unmarshal(data []byte, v interface{}) error

		UnmarshalFromString(str string, v interface{}) error
	}

	// Response 建议实现的结构体的导出字段加上json tag
	Response interface {
		GetCode() string

		SetCode(code string)

		GetMessage() string

		SetMessage(msg string)

		GetData() interface{}

		SetData(data interface{})
	}

	RespProvider func(code, msg string, data interface{}) Response
)
