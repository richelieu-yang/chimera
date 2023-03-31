package ginKit

import (
	"github.com/richelieu42/chimera/src/http/refererKit"
)

type (
	Config struct {
		Host string `json:"host,optional"`
		Port int    `json:"port,optional"`
		/*
			日志的颜色（默认true）
			true: 	强制设置日志颜色
			false: 	禁止日志颜色
		*/
		Colorful   bool `json:"colorful,default=true"`
		Middleware *MiddlewareConfig
		Ssl        *SslConfig `json:"ssl"`
	}

	MiddlewareConfig struct {
		Gzip          bool                                 `json:"gzip,default=true"`
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
		Port     int    `json:"port,optional"`
	}
)
