package wsKit

import (
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGinHandler(t *testing.T) {
	engine := gin.Default()

	engine.Any("/ping", GinHandler)

	if err := engine.Run(":8080"); err != nil {
		panic(err)
	}
}
