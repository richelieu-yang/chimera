package ginKit

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// NewOpenCorsMiddleware 新建一个开放的（任意origin都通过）cors中间件.
func NewOpenCorsMiddleware() gin.HandlerFunc {
	return NewCorsMiddleware(nil)
}

// NewCorsMiddleware 新建一个cors中间件.
/*
@param origins origin白名单，可以为nil
*/
func NewCorsMiddleware(origins []string) gin.HandlerFunc {
	config := newCorsConfig(origins)
	return cors.New(config)
}

/*
@param origins origin白名单，可以为nil（任意origin都通过）
@return cors依赖的配置
*/
func newCorsConfig(origins []string) cors.Config {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Last-Modified"},
		MaxAge:           time.Hour * 12,
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
