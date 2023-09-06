package pulsarKit

type (
	Config struct {
		// Addresses Broker地址
		Addresses []string `json:"addresses" yaml:"addresses"`

		VerifyConfig VerifyConfig `json:"verify,optional" yaml:"verify"`
	}

	VerifyConfig struct {
		// Topic 用于验证"pulsar服务是否正常启动"的topic
		Topic string `json:"topic,optional" yaml:"topic"`
		// Print 是否输出 验证日志 到控制台？
		Print bool `json:"print,default=false" yaml:"print"`
	}
)
