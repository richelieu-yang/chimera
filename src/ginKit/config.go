package ginKit

import (
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/http/refererKit"
)

type (
	Config struct {
		Host string `json:"host,optional"`
		Port int    `json:"port,default=-1,range=[-1:65535]"`
		/*
			日志的颜色（默认true）
			true: 	强制设置日志颜色
			false: 	禁止日志颜色
		*/
		Colorful   bool              `json:"colorful,default=true"`
		Middleware *MiddlewareConfig `json:"middleware,optional"`
		SSL        *SslConfig        `json:"ssl,optional"`
	}

	MiddlewareConfig struct {
		Gzip          bool                                 `json:"gzip,default=false"`
		XFrameOptions string                               `json:"xFrameOptions,optional"`
		Cors          *CorsConfig                          `json:"cors,optional"`
		Referer       []*refererKit.RefererVerifierBuilder `json:"referer,optional"`
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		Origins []string `json:"origins,optional"`
	}

	SslConfig struct {
		CertFile string `json:"certFile"`
		KeyFile  string `json:"keyFile"`
		Port     int    `json:"port,default=-1,range=[-1:65535]"`
	}
)

func (config *Config) Check() error {
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	// ssl
	sslConfig := config.SSL
	if sslConfig != nil {
		if sslConfig.Port == -1 || strKit.HasEmpty(sslConfig.CertFile, sslConfig.KeyFile) {
			sslConfig = nil
		} else {
			if err := fileKit.AssertExistAndIsFile(sslConfig.CertFile); err != nil {
				return err
			}
			if err := fileKit.AssertExistAndIsFile(sslConfig.KeyFile); err != nil {
				return err
			}
		}
	}

	// http port
	if config.Port != -1 {
		if sslConfig != nil && config.Port == sslConfig.Port {
			return errorKit.Simple("http port and https port are same(%d)", config.Port)
		}
	} else {
		if sslConfig == nil {
			return errorKit.Simple("both http port and https port are invalid")
		}
	}

	return nil
}
