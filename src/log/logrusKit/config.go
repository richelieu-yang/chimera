package logrusKit

type (
	Config struct {
		Level string `json:"level" yaml:"level"`

		PrintBasic bool `json:"printBasic" yaml:"printBasic"`

		// DisableQuote
		/*
			true: 	禁用 日志内容的双引号.
			false: 	允许 日志内容的双引号.
		*/
		DisableQuote bool `json:"disableQuote" yaml:"disableQuote"`
	}
)
