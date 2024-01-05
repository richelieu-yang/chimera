package logrusKit

type (
	Config struct {
		// Level 日志级别，默认: debug
		Level string `json:"level" yaml:"level"`

		// EnableQuote 是否允许日志内容的双引号?
		EnableQuote bool `json:"enableQuote" yaml:"enableQuote"`

		// PrintBasic 是否打印基本信息？
		PrintBasic bool `json:"printBasic" yaml:"printBasic"`
	}
)
