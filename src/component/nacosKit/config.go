package nacosKit

import "github.com/nacos-group/nacos-sdk-go/v2/common/constant"

type (
	Config struct {
		/* client */
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

func a() {
	constant.ServerConfig
}
