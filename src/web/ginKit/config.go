package ginKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/web/refererKit"
)

type (
	Config struct {
		HostName string `json:"hostName,optional"`
		Port     int    `json:"port,default=80,range=[1:65535]"`
		Colorful bool   `json:"colorful,default=true"`
		Pprof    bool   `json:"pprof,default=false"`

		SSL        SslConfig        `json:"ssl"`
		Middleware MiddlewareConfig `json:"middleware"`
	}

	SslConfig struct {
		Access   bool   `json:"access,default=false"`
		CertFile string `json:"certFile,optional"`
		KeyFile  string `json:"keyFile,optional"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64                                `json:"bodyLimit,default=-1,range=[-1:]"`
		Gzip          bool                                 `json:"gzip,default=false"`
		XFrameOptions string                               `json:"xFrameOptions,optional"`
		Cors          CorsConfig                           `json:"cors,optional"`
		Referer       []*refererKit.RefererVerifierBuilder `json:"referer,optional"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access,default=false"`
		Origins []string `json:"origins,optional"`
	}
)

func (config *Config) Verify() error {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	// ssl
	sslConfig := config.SSL
	if sslConfig.Access {
		if err := fileKit.AssertExistAndIsFile(sslConfig.CertFile); err != nil {
			return err
		}
		if err := fileKit.AssertExistAndIsFile(sslConfig.KeyFile); err != nil {
			return err
		}
	}

	return nil
}
