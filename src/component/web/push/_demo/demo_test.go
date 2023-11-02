package _demo

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/sseKit"
	"github.com/richelieu-yang/chimera/v2/src/goroutine/poolKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestDemo(t *testing.T) {
	logrusKit.MustSetUp(nil)

	pool, err := poolKit.NewPool(2000)
	if err != nil {
		logrus.Fatal(err)
	}
	pushKit.MustSetUp(pool)

	engine := gin.Default()
	engine.Use(ginKit.NewCorsMiddleware(nil))

	// SSE
	{
		processor, err := sseKit.NewProcessor(nil, &demoListener{}, sseKit.MessageTypeRaw)
		if err != nil {
			panic(err)
		}
		engine.GET("/sse", processor.ProcessWithGin)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
