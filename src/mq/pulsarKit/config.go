package pulsarKit

type (
	Config struct {
		// Addresses Broker地址
		Addresses []string

		// LogDir 文件日志路径（默认: 输出到控制台）
		LogDir string

		// TopicForVerify 用于验证"pulsar服务是否正常启动"的topic（为空则不测试）
		TopicForVerify string
	}
)
