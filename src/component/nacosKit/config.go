package nacosKit

type (
	Addr string

	Config struct {
		Server []*ServerConfig
	}

	ServerConfig struct {
		// IpAddr the nacos server address
		IpAddr string `json:"ip" yaml:"ip"`
		// Port nacos server port
		Port uint64 `json:"ip" yaml:"ip"`
	}
)
