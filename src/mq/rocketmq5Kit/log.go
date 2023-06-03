package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/core/osKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/operationKit"
	"os"
	"sync"
)

type (
	// LogConfig 如果输出到日志文件且该文件已存在，会"追加到最后面".
	LogConfig struct {
		// ToConsole RocketMQ5的客户端日志是否输出到控制台？
		ToConsole bool
		// LogDir RocketMQ5的客户端日志要输出到文件的情况下，日志所在的目录
		LogDir string
		// LogName RocketMQ5的客户端日志要输出到文件的情况下，日志的文件名
		LogName string
	}
)

var (
	// 主要是 客户端输出（无论是到控制台还是文件） 涉及到修改环境变量，目前只能靠 加锁 解决
	lock = new(sync.Mutex)

	defaultLogConfig = &LogConfig{
		ToConsole: true,
		LogDir:    "",
		LogName:   "",
	}
)

// SetLogout
/*
!!!: 如果是输出到文件的话，由于使用的是系统变量，需要考虑"并发问题".
*/
func (lc *LogConfig) SetLogout() error {
	c := operationKit.Ternary(lc != nil, lc, defaultLogConfig)

	if c.ToConsole {
		// (1) 输出到控制台
		if err := os.Setenv(rmq_client.ENABLE_CONSOLE_APPENDER, "true"); err != nil {
			return err
		}
	} else {
		// (2) 输出到日志文件
		logDir := c.LogDir
		if err := fileKit.MkDirs(logDir); err != nil {
			return err
		}
		logDir = strKit.EmptyToDefault(logDir, ".")

		logName := c.LogName
		if strKit.IsEmpty(logName) {
			return errorKit.New("logName is empty")
		}

		if err := osKit.SetEnvs(map[string]string{
			rmq_client.ENABLE_CONSOLE_APPENDER: "false",
			rmq_client.CLIENT_LOG_ROOT:         logDir,
			rmq_client.CLIENT_LOG_FILENAME:     logName,
		}); err != nil {
			return err
		}
	}

	rmq_client.ResetLogger()
	return nil
}
