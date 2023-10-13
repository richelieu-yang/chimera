package ginKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/refererKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

type (
	Config struct {
		Mode     string `json:"mode,default=debug,options=debug|release|test" yaml:"mode"`
		HostName string `json:"hostName,optional" yaml:"hostName"`
		// Port
		/*
			-1（默认）: 不使用 http port
		*/
		Port     int  `json:"port,default=-1,range=[-1:65535]" yaml:"port"`
		Colorful bool `json:"colorful,default=true" yaml:"colorful"`
		Pprof    bool `json:"pprof,default=false" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware"`
	}

	SslConfig struct {
		// Port
		/*
			-1（默认）: 不使用 https port
		*/
		Port     int    `json:"port,default=-1,range=[-1:65535]" yaml:"port"`
		CertFile string `json:"certFile,optional" yaml:"certFile"`
		KeyFile  string `json:"keyFile,optional" yaml:"keyFile"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64                                `json:"bodyLimit,default=-1,range=[-1:]" yaml:"bodyLimit"`
		Gzip          bool                                 `json:"gzip,default=false" yaml:"gzip"`
		XFrameOptions string                               `json:"xFrameOptions,optional" yaml:"xFrameOptions"`
		Cors          CorsConfig                           `json:"cors" yaml:"cors"`
		Referer       []*refererKit.RefererVerifierBuilder `json:"referer,optional" yaml:"referer"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access,default=false" yaml:"access"`
		Origins []string `json:"origins,optional" yaml:"origins"`
	}
)

func (config *Config) Verify() error {
	// https server
	ssl := config.SSL
	if ssl.Port != -1 {
		if err := netKit.AssertValidPort(ssl.Port); err != nil {
			return errorKit.Wrap(err, "invalid https port(%d)", ssl.Port)
		}
		if err := fileKit.AssertExistAndIsFile(ssl.CertFile); err != nil {
			return err
		}
		if err := fileKit.AssertExistAndIsFile(ssl.KeyFile); err != nil {
			return err
		}
	}

	return nil
}
