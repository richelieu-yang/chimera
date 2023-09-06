package ginKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/web/refererKit"
)

type (
	Config struct {
		Mode     string `json:"mode,default=debug,options=debug|release|test" yaml:"mode"`
		HostName string `json:"hostName,optional" yaml:"hostName"`
		Port     int    `json:"port,default=80,range=[1:65535]" yaml:"port"`
		Colorful bool   `json:"colorful,default=true" yaml:"colorful"`
		Pprof    bool   `json:"pprof,default=false" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware"`
	}

	SslConfig struct {
		Access   bool   `json:"access,default=false" yaml:"access"`
		CertFile string `json:"certFile,optional" yaml:"certFile"`
		KeyFile  string `json:"keyFile,optional" yaml:"keyFile"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64                                `json:"bodyLimit,default=-1,range=[-1:]" yaml:"bodyLimit"`
		Gzip          bool                                 `json:"gzip,default=false" yaml:"gzip"`
		XFrameOptions string                               `json:"xFrameOptions,optional" yaml:"xFrameOptions"`
		Cors          CorsConfig                           `json:"cors"`
		Referer       []*refererKit.RefererVerifierBuilder `json:"referer,optional"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access,default=false" yaml:"access"`
		Origins []string `json:"origins,optional" yaml:"origins"`
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
