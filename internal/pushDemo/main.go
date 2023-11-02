package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v2/internal/pushDemo/docs"
	"github.com/richelieu-yang/chimera/v2/internal/pushDemo/handler"
	"github.com/richelieu-yang/chimera/v2/internal/pushDemo/types"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/sseKit"
	"github.com/richelieu-yang/chimera/v2/src/goroutine/poolKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type resp struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	logrusKit.MustSetUp(nil)
	jsonRespKit.MustSetUp(func(code, msg string, data interface{}) interface{} {
		return &resp{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	})

	pool, err := poolKit.NewPool(2000)
	if err != nil {
		logrus.Fatal(err)
	}
	pushKit.MustSetUp(pool)

	engine := gin.Default()
	engine.Use(ginKit.NewCorsMiddleware(nil))

	// push
	{
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		engine.POST("/push_to_all", ginKit.WrapToHandlerFunc(handler.PushToAll))
		engine.POST("/push_to_bsid", ginKit.WrapToHandlerFunc(handler.PushToBsid))
		engine.POST("/push_to_user", ginKit.WrapToHandlerFunc(handler.PushToUser))
		engine.POST("/push_to_group", ginKit.WrapToHandlerFunc(handler.PushToGroup))
	}

	// SSE
	{
		processor, err := sseKit.NewProcessor(nil, &types.DemoListener{}, sseKit.MessageTypeRaw)
		if err != nil {
			panic(err)
		}
		engine.GET("/sse", processor.ProcessWithGin)
	}

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
