// Package nacosKit
/*
官方文档（英文）：https://github.com/nacos-group/nacos-sdk-go
官方文档（中文）：https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md
*/
package nacosKit

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/mainControl"
	"github.com/richelieu42/go-scales/src/netKit"
)

// NewNamingClient 服务发现客户端
/*
使用方法详见：https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md#%E6%9C%8D%E5%8A%A1%E5%8F%91%E7%8E%B0

@param logDir 	日志的目录（会尝试创建目录）
@param level 	日志级别（可以为nil，默认依据 mainControl.IsDebug()）
*/
func NewNamingClient(config *NacosConfig, outputDir string, logLevel NacosLogLevel) (naming_client.INamingClient, error) {
	clientConfig, serverConfigs, err := getConfigs(config, outputDir, logLevel)
	if err != nil {
		return nil, err
	}
	return clients.NewNamingClient(vo.NacosClientParam{
		ClientConfig:  clientConfig,
		ServerConfigs: serverConfigs,
	})
}

// NewConfigClient 动态配置客户端
/*
使用方法详见：https://github.com/nacos-group/nacos-sdk-go/blob/master/README_CN.md#%E5%8A%A8%E6%80%81%E9%85%8D%E7%BD%AE
*/
func NewConfigClient(config *NacosConfig, outputDir string, logLevel NacosLogLevel) (config_client.IConfigClient, error) {
	clientConfig, serverConfigs, err := getConfigs(config, outputDir, logLevel)
	if err != nil {
		return nil, err
	}
	return clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  clientConfig,
		ServerConfigs: serverConfigs,
	})
}

func getConfigs(config *NacosConfig, outputDir string, logLevel NacosLogLevel) (*constant.ClientConfig, []constant.ServerConfig, error) {
	if config == nil {
		return nil, nil, errorKit.Simple("config is nil")
	}
	// outputDir
	if err := fileKit.MkDirs(outputDir); err != nil {
		return nil, nil, err
	}
	// logLevel
	if logLevel == nil {
		if mainControl.IsDebug() {
			logLevel = DebugLevel
		} else {
			logLevel = InfoLevel
		}
	}

	// clientConfig
	clientConfig, err := newClientConfig(config, outputDir, logLevel)
	if err != nil {
		return nil, nil, err
	}
	// serverConfigs
	serverConfigs, err := newServerConfigs(config)
	if err != nil {
		return nil, nil, err
	}
	return &clientConfig, serverConfigs, nil
}

func newClientConfig(config *NacosConfig, outputDir string, level NacosLogLevel) (constant.ClientConfig, error) {
	logDir := pathKit.Join(outputDir, "log")
	if err := fileKit.MkDirs(logDir); err != nil {
		return constant.ClientConfig{}, err
	}
	cacheDir := pathKit.Join(outputDir, "cache")
	if err := fileKit.MkDirs(cacheDir); err != nil {
		return constant.ClientConfig{}, err
	}

	return constant.ClientConfig{
		// 请求Nacos服务端的超时时间，默认是10000ms
		TimeoutMs: 10000,
		// ACM的命名空间Id
		NamespaceId: config.NamespaceId,

		// 缓存service信息的目录，默认是当前运行目录
		CacheDir: cacheDir,
		// 在启动的时候不读取缓存在CacheDir的service信息
		NotLoadCacheAtStart: true,

		// 日志存储路径
		LogDir: logDir,
		// 日志默认级别，值必须是：debug,info,warn,error，默认值是info
		LogLevel: *level,

		/* Richelieu: 由于依赖更新，下面两个属性被移除 */
		//// 日志轮转周期，比如：30m, 1h, 24h, 默认是24h
		//RotateTime: "24h",
		//// 日志最大文件数，默认3
		//MaxAge: 3,
	}, nil
}

// newServerConfigs
/*
@return 第一个返回值"非nil"且"长度>0"
*/
func newServerConfigs(config *NacosConfig) ([]constant.ServerConfig, error) {
	addresses := config.ServerAddresses
	length := len(addresses)
	if length == 0 {
		return nil, errorKit.Simple("length of addresses is 0")
	}

	configs := make([]constant.ServerConfig, 0, length)
	for _, addr := range addresses {
		info, err := netKit.ParseStringToAddress(addr)
		if err != nil {
			return nil, err
		}
		configs = append(configs, constant.ServerConfig{
			IpAddr:      info.Host,
			Port:        uint64(info.Port),
			ContextPath: "/nacos",
			Scheme:      "http",
		})
	}
	return configs, nil
}
