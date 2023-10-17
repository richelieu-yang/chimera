package pulsarKit

type (
	Config struct {
		// Addrs Broker地址
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required,gte=1,dive,hostname_port"`
	}

	//VerifyConfig struct {
	//	// Topic 用于验证"pulsar服务是否正常启动"的topic
	//	Topic string `json:"topic" yaml:"topic"`
	//	// Print 是否输出 验证日志 到控制台？
	//	Print bool `json:"print" yaml:"print"`
	//}
)
