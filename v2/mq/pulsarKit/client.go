package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu42/chimera/v2/core/sliceKit"
	"github.com/richelieu42/chimera/v2/core/strKit"
	"github.com/richelieu42/chimera/v2/log/logrusKit"
	"github.com/richelieu42/chimera/v2/netKit"
	"github.com/sirupsen/logrus"
)

func NewClient(addresses []string, logPath string) (pulsar.Client, error) {
	/* url */
	tmp, err := netKit.ProcessAddresses(addresses)
	if err != nil {
		return nil, err
	}
	url := UrlPrefix + sliceKit.Join(tmp, ",")

	/* logger */
	var logger log.Logger
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

	return client, nil
}
