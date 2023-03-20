package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/apache/pulsar-client-go/pulsar/log"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/netKit"
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

	var logger log.Logger
	//logPath := strKit.Trim(config.LogPath)
	//if strKit.IsNotEmpty(logPath) {
	//	if err := fileKit.MkParentDirs(logPath); err != nil {
	//		return nil, err
	//	}
	//	log.NewLoggerWithLogrus()
	//}

	return pulsar.NewClient(pulsar.ClientOptions{
		URL:    url,
		Logger: logger,
	})
}
