package componentKit

import (
	"fmt"
	"github.com/richelieu42/chimera/src/confKit"
	"github.com/richelieu42/chimera/src/consts"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/richelieu42/chimera/src/database/redisKit"
	"github.com/richelieu42/chimera/src/ginKit"
	"github.com/richelieu42/chimera/src/jsonKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
	"time"
)

type (
	EnvConfig struct {
		Ip string

		Runtime *RuntimeConfig
		Logrus  *logrusKit.Config
		Gin     *ginKit.GinConfig
		Redis   *redisKit.Config

		RocketMQ5 *rocketmq5Kit.Config
	}
)

var (
	// defaultEnvMap 默认值
	defaultEnvMap = map[string]interface{}{
		// logrus
		"logrus.level":           "debug",
		"logrus.timestampFormat": string(timeKit.FormatEntire),

		// gin
		"gin.port":            80,
		"gin.colorful":        true,
		"gin.middleware.gzip": true,

		// redis
		"redis.mode": -1,

		// RocketMQ5
		"rocketmq5.credentials.accessKey":    "",
		"rocketmq5.credentials.accessSecret": "",
	}
	EnvNotLoadedError     = errorKit.Simple("env file hasn't been loaded")
	EnvAlreadyLoadedError = errorKit.Simple("env file has already been loaded")
)

// 服务的启动时间
var startingTime = time.Now()

var envConfig *EnvConfig = nil

func GetIp() (string, error) {
	if envConfig == nil {
		return "", EnvNotLoadedError
	}
	return envConfig.Ip, nil
}

// GetAddress 获取当前服务的地址
func GetAddress() (string, error) {
	ip, err := GetIp()
	if err != nil {
		return "", err
	}
	ginConfig, err := GetGinConfig()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s:%d", ip, ginConfig.Port), nil
}

// InitializeEnvironment 加载 env.yaml; 初始化组件: logrus、runtime
/*
@param pathArgs 第1个值代表（如果有的话）: env.yaml的相对路径
*/
func InitializeEnvironment() error {
	if envConfig != nil {
		// 理论上此方法只会被调用1次，因此此处会返回error
		return EnvAlreadyLoadedError
	}

	// 先简单初始化下logrus，下面会完整地初始化logrus组件
	logrusKit.InitializeByDefault()

	if err := loadEnvYaml(); err != nil {
		return err
	}

	// (1) 初始化 logrus
	if err := initializeLogrusComponent(); err != nil {
		return err
	}
	// (2) 初始化 runtime
	if err := initializeRuntimeComponent(); err != nil {
		return err
	}
	return nil
}

// loadEnvYaml 加载配置文件 env.yaml
func loadEnvYaml() error {
	absPath := pathKit.Join(pathKit.GetProjectDir(), consts.EnvPath)

	envConfig = &EnvConfig{}
	if err := confKit.ReadFileAs(absPath, defaultEnvMap, envConfig); err != nil {
		envConfig = nil
		return err
	}

	/* 优化配置文件 */
	envConfig.Redis.Simplify()

	logrusKit.Initialize(logrus.DebugLevel, timeKit.FormatEntire)
	json, _ := jsonKit.MarshalWithIndent(envConfig)
	logrus.Infof("[COMPONENT, ENV] environment:\n%s", json)

	logrus.Info("[COMPONENT, ENV] Initialize successfully.")
	return nil
}

func GetStartingTime() time.Time {
	return startingTime
}

// GetRuntimeConfig
/*
@return 有可能2个返回值都为nil
*/
func GetRuntimeConfig() (*RuntimeConfig, error) {
	if envConfig == nil {
		return nil, EnvNotLoadedError
	}
	return envConfig.Runtime, nil
}

// GetLogrusConfig
/*
@return 两个返回值中必定有一个不为nil（因为有默认值）
*/
func GetLogrusConfig() (*logrusKit.Config, error) {
	if envConfig == nil {
		return nil, EnvNotLoadedError
	}
	return envConfig.Logrus, nil
}

// GetGinConfig
/*
@return 两个返回值中必定有一个不为nil（因为有默认值）
*/
func GetGinConfig() (*ginKit.GinConfig, error) {
	if envConfig == nil {
		return nil, EnvNotLoadedError
	}
	return envConfig.Gin, nil
}

// GetRocketmq5Config
/*
@return 两个返回值有可能都为nil（因为没有默认值）
*/
func GetRocketmq5Config() (*rocketmq5Kit.Config, error) {
	if envConfig == nil {
		return nil, EnvNotLoadedError
	}
	return envConfig.RocketMQ5, nil
}
