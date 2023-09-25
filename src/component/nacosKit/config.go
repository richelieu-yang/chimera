package nacosKit

type (
	Config struct {
		/* client */
		// NamespaceId 命名空间的id，配置 "" 和 "public" 效果一样（都是使用保留空间public）
		NamespaceId string `json:"namespaceId" yaml:"namespaceId"`

		/* server */
		Addresses []string `json:"addresses" yaml:"addresses"`
	}

	ServerConfig struct {
		// the nacos server scheme,default=http,this is not required in 2.0
		Scheme string `json:"scheme" yaml:"scheme"`
		// the nacos server contextpath,default=/nacos,this is not required in 2.0
		ContextPath string `json:"contextPath" yaml:"contextPath"`
		// the nacos server address
		IpAddr string `json:"ipAddr" yaml:"ipAddr"`
		// nacos server port
		Port uint64 `json:"port" yaml:"port"`

		//// nacos server grpc port, default=server port + 1000, this is not required
		//GrpcPort uint64
	}
)
