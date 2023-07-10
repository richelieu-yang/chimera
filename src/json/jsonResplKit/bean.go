package jsonResplKit

type (
	API interface {
		//Marshal(v interface{}) ([]byte, error)

		MarshalToString(v interface{}) (string, error)

		//Unmarshal(data []byte, v interface{}) error
		//
		//UnmarshalFromString(str string, v interface{}) error
	}

	// BeanProvider
	/*
		@return 建议是个结构体实例 && 加上json tag
	*/
	BeanProvider func(code, msg string, data interface{}) interface{}
)
