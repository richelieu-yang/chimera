package i18nPackKit

type (
	// Maker
	/*
		@return (1) 返回值是一个结构体实例指针
				(2) 结构体建议加上json tag
	*/
	Maker func(code, msg string, data interface{}) interface{}

	bean struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}
)
