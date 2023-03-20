package pulsarKit

type (
	Config struct {
		// Addresses Broker地址
		Addresses []string
		// LogPath 文件日志路径（默认: 输出到控制台）
		LogPath string
		// TopicForTest 用于测试连通的topic（为空则不测试）
		TopicForTest string
	}
)
