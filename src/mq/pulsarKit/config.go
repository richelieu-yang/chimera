package pulsarKit

type (
	Config struct {
		// Addresses Broker地址
		Addresses []string
		// LogDir 日志目录（默认: 输出到控制台）
		LogDir string
		// TopicForTest 用于测试连通的topic（为空则不测试）
		TopicForTest string
	}
)
