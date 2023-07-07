package jsonResplKit

type (
	Response interface {
		GetCode() string

		SetCode(code string)

		GetMessage() string

		SetMessage(msg string)

		GetData() interface{}

		SetData(data interface{})
	}

	ResponseProvider func(code, msg string, data interface{}) Response
)
