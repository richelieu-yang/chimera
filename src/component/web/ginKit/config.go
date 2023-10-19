package ginKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/refererKit"
)

type (
	Config struct {
		Mode     string `json:"mode,default=debug,options=debug|release|test" yaml:"mode"`
		HostName string `json:"hostName" yaml:"hostName"`
		// Port
		/*
			0（默认）: 不使用 http port
		*/
		Port     int  `json:"port" yaml:"port"`
		Colorful bool `json:"colorful" yaml:"colorful"`
		Pprof    bool `json:"pprof" yaml:"pprof"`

		SSL        SslConfig        `json:"ssl" yaml:"ssl"`
		Middleware MiddlewareConfig `json:"middleware" yaml:"middleware"`
	}

	SslConfig struct {
		// Port
		/*
			0（默认）: 不使用 https port
		*/
		Port     int    `json:"port" yaml:"port"`
		CertFile string `json:"certFile" yaml:"certFile"`
		KeyFile  string `json:"keyFile" yaml:"keyFile"`
	}

	MiddlewareConfig struct {
		BodyLimit     int64                                `json:"bodyLimit" yaml:"bodyLimit"`
		Gzip          bool                                 `json:"gzip" yaml:"gzip"`
		XFrameOptions string                               `json:"xFrameOptions" yaml:"xFrameOptions"`
		Cors          CorsConfig                           `json:"cors" yaml:"cors"`
		Referer       []*refererKit.RefererVerifierBuilder `json:"referer" yaml:"referer"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Access  bool     `json:"access" yaml:"access"`
		Origins []string `json:"origins" yaml:"origins"`
	}
)

//func (config *Config) Verify() error {
//	// https server
//	ssl := config.SSL
//	if ssl.Port != -1 {
//		if err := netKit.AssertValidPort(ssl.Port); err != nil {
//			return errorKit.Wrap(err, "invalid https port(%d)", ssl.Port)
//		}
//		if err := fileKit.AssertExistAndIsFile(ssl.CertFile); err != nil {
//			return err
//		}
//		if err := fileKit.AssertExistAndIsFile(ssl.KeyFile); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
