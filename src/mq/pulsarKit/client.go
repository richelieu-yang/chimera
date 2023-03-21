package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu42/chimera/src/core/errorKit"
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

	/* url */
	tmp, err := netKit.ProcessAddresses(config.Addresses)
	if err != nil {
		return nil, err
	}
	url := UrlPrefix + sliceKit.Join(tmp, ",")

	/* logger */
	var logger log.Logger
	logPath := config.LogPath
	if strKit.IsNotEmpty(logPath) {
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
