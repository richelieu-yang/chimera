package logrusKit

type (
	Config struct {
		Level      string `json:"level,default=debug" yaml:"level"`
		PrintBasic bool   `json:"printBasic,default=false" yaml:"printBasic"`
	}
)
