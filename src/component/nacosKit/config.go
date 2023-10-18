package nacosKit

type (
	Config struct {
		/* client */
		// NamespaceId 命名空间的id，配置 "" 和 "public" 效果一样（都是使用保留空间public）
		NamespaceId string `json:"namespaceId" yaml:"namespaceId"`

		/* server */
		Addrs []string `json:"addrs" yaml:"addrs" validate:"required"`
	}
)
