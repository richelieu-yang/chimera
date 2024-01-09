package ginKit

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func NewFastGzipMiddleware() gin.HandlerFunc {
	return gzip.Gzip(gzip.BestSpeed)
}
