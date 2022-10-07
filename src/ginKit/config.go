package ginKit

import "github.com/richelieu42/go-scales/src/netKit"

type (
	GinConfig struct {
		Host string
		Port int
		/*
			日志的颜色（默认true）
			true: 	强制设置日志颜色
			false: 	禁止日志颜色
		*/
		Colorful   bool
		Middleware *MiddlewareConfig
	}

	MiddlewareConfig struct {
		Gzip          bool
		XFrameOptions string
		Cors          *CorsConfig
	}

	// CorsConfig cors（跨源资源共享）的配置
	CorsConfig struct {
		// 是否对cors进行配置？
		Access  bool
		Origins []string
	}
)

func (gc *GinConfig) GetAddress() string {
	return netKit.JoinHostPort(gc.Host, gc.Port)
}
