package pulsarKit

type (
	Config struct {
		// Addresses Broker地址
		Addresses []string `json:"addresses"`

		VerifyConfig *VerifyConfig `json:"_verify"`
	}

	VerifyConfig struct {
		// Topic 用于验证"pulsar服务是否正常启动"的topic
		Topic string `json:"topic"`
		// Print 是否输出 验证日志 到控制台？
		Print bool `json:"print"`
	}
)
