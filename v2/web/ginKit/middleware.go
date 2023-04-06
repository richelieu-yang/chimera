package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/core/sliceKit"
	"github.com/richelieu42/chimera/v2/core/strKit"
	"github.com/richelieu42/chimera/v2/web/refererKit"
)

// AttachCommonMiddlewares 绑定一些常用的中间件.
func AttachCommonMiddlewares(engine *gin.Engine, middlewareConfig *MiddlewareConfig, recoveryMiddleware gin.HandlerFunc) error {
	// gzip
	/*
		PS:
		(1) 必须在 recoveryMiddleware 前面，为了万无一失还是放在最前面吧；
		(2) gzip会使得响应头中的 Content-Length 不生效.
	*/
	if middlewareConfig != nil && middlewareConfig.Gzip {
		engine.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// logger(necessary)
	engine.Use(gin.Logger())

	// recovery(necessary)
	if recoveryMiddleware == nil {
		recoveryMiddleware = gin.Recovery()
	}
	engine.Use(recoveryMiddleware)

	if middlewareConfig != nil {
		// cors
		{
			var origins []string
			corsConfig := middlewareConfig.Cors
			if corsConfig != nil {
				origins = corsConfig.Origins
			}
			origins = sliceKit.RemoveEmpty(origins, true)
			origins = sliceKit.Uniq(origins)

			engine.Use(NewCorsMiddleware(origins))
		}

		// referer（必须在cors中间件后面）
		{
			refererConfig := middlewareConfig.Referer
			if refererConfig != nil {
				middleware, err := refererKit.NewGinRefererMiddleware(refererConfig)
				if err != nil {
					return err
				}
				engine.Use(middleware)
			}
		}

		// others
		engine.Use(func(ctx *gin.Context) {
			if strKit.IsNotEmpty(middlewareConfig.XFrameOptions) {
				// e.g.不能被嵌入到任何iframe或frame中
				ctx.Header("X-Frame-Options", middlewareConfig.XFrameOptions)
			}
			// 解决漏洞: 未启用Web浏览器XSS保护
			ctx.Header("X-XSS-Protection", "1;mode=block")
		})
	}

	return nil
}
