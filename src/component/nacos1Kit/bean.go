package nacos1Kit

type (
	NacosConfig struct {
		ServerAddresses []string
		NamespaceId     string
		Naming          NacosNamingConfig
	}

	// NacosNamingConfig "服务发现"的配置
	NacosNamingConfig struct {
		// 注册的服务实例是否是临时的？
		Ephemeral bool
	}
)
