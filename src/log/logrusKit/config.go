package logrusKit

type (
	Config struct {
		Level      string `json:"level,default=debug"`
		PrintBasic bool   `json:"printBasic,default=false"`
	}
)
