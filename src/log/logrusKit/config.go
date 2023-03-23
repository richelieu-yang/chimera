package logrusKit

type (
	Config struct {
		Level      string `json:"level"`
		PrintBasic bool   `json:"printBasic"`
	}
)
