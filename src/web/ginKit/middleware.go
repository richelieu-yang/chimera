package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/web/refererKit"
	"net/http"
)

// attachMiddlewares 绑定一些常用的中间件.
func attachMiddlewares(engine *gin.Engine, config *MiddlewareConfig, recoveryMiddleware gin.HandlerFunc) error {
	// gzip
	/*
		PS:
		(1) 必须在 recoveryMiddleware 前面，为了万无一失还是放在最前面吧；
		(2) gzip会使得响应头中的 Content-Length 不生效.
	*/
	if config != nil && config.Gzip {
		engine.Use(gzip.Gzip(gzip.BestSpeed))
	}

	// logger(necessary)
	engine.Use(gin.Logger())

	// recovery(necessary)
	if recoveryMiddleware == nil {
		recoveryMiddleware = gin.Recovery()
	}
	engine.Use(recoveryMiddleware)

	if config != nil {
		// cors
		{
			var origins []string
			corsConfig := config.Cors
			if corsConfig != nil {
				origins = corsConfig.Origins
			}
			origins = sliceKit.RemoveEmpty(origins, true)
			origins = sliceKit.Uniq(origins)

			engine.Use(NewCorsMiddleware(origins))
		}

		// referer（必须在cors中间件后面）
		{
			refererConfig := config.Referer
			if refererConfig != nil {
				middleware, err := refererKit.NewGinRefererMiddleware(refererConfig)
				if err != nil {
					return err
				}
				engine.Use(middleware)
			}
		}

		// bodyLimit
		if config.BodyLimit > 0 {
			engine.Use(func(ctx *gin.Context) {
				if ctx.Request.Body != nil {
					ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, config.BodyLimit<<20+256)
				}
			})
		}

		// others
		engine.Use(func(ctx *gin.Context) {
			if strKit.IsNotEmpty(config.XFrameOptions) {
				// e.g.不能被嵌入到任何iframe或frame中
				ctx.Header("X-Frame-Options", config.XFrameOptions)
			}
			// 解决漏洞: 未启用Web浏览器XSS保护
			ctx.Header("X-XSS-Protection", "1;mode=block")
		})
	}

	return nil
}
