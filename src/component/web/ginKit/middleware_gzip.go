package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

// NewFastGzipMiddleware
/*
PS: 涉及多个服务（请求转发）的场景下，(1) 最外层的务使用gzip压缩;
								(2) 内层的服务不使用gzip压缩.
*/
func NewFastGzipMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.BestSpeed)
}
