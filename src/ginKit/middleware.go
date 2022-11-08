package ginKit

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/http/refererKit"
	"time"
)

// AttachCommonMiddlewares 绑定一些常用的中间件
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
	if recoveryMiddleware != nil {
		engine.Use(recoveryMiddleware)
	} else {
		engine.Use(gin.Recovery())
	}

	if middlewareConfig != nil {
		// referer
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

		// cors
		{
			corsConfig := middlewareConfig.Cors
			if corsConfig != nil {
				origins := corsConfig.Origins
				// 去除无效项
				origins = sliceKit.RemoveEmpty(origins, true)
				if len(origins) >= 0 {
					engine.Use(cors.New(newCorsConfig(origins)))
				}
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

/*
@param origins origin白名单，可以为nil
@return cors依赖的配置
*/
func newCorsConfig(origins []string) cors.Config {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Last-Modified"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
		AllowWebSockets:  true,
		AllowWildcard:    true,
	}
	if len(origins) > 0 {
		// 允许部分
		config.AllowOrigins = origins
	} else {
		// 允许全部
		config.AllowOriginFunc = func(origin string) bool {
			return true
		}
	}
	return config
}
