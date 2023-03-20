package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/netKit"
	"github.com/sirupsen/logrus"
)

func NewClient(config *Config) (pulsar.Client, error) {
	if config == nil {
		return nil, errorKit.Simple("config == nil")
	}

	tmp, err := netKit.ProcessAddresses(config.Addresses)
	if err != nil {
		return nil, err
	}
	url := UrlPrefix + sliceKit.Join(tmp, ",")

	// pulsar客户端的日志输出
	var logger log.Logger
	logDir := strKit.Trim(config.LogDir)
	if strKit.IsNotEmpty(logDir) {
		if err := fileKit.MkParentDirs(logDir); err != nil {
			return nil, err
		}
		logPath := pathKit.Join(logDir, ClientLogName)
		fileLogger, err := logrusKit.NewFileLogger(logPath, nil, logrus.DebugLevel, false)
		if err != nil {
			return nil, err
		}
		logger = log.NewLoggerWithLogrus(fileLogger)
	}

	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL:    url,
		Logger: logger,
	})
	if err != nil {
		return nil, err
	}

	topic := config.TopicForVerify
	if strKit.IsNotEmpty(topic) {
		if err := VerifyPulsar(client, topic); err != nil {
			return nil, err
		}
	}

	return client, nil
}
