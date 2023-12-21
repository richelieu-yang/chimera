package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/richelieu-yang/chimera/v2/internal/pushDemo/docs"
	"github.com/richelieu-yang/chimera/v2/internal/pushDemo/handler"
	"github.com/richelieu-yang/chimera/v2/internal/pushDemo/types"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/sseKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/wsKit"
	"github.com/richelieu-yang/chimera/v2/src/concurrency/poolKit"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func init() {
	cpuKit.SetUp()
	logrusKit.MustSetUp(nil)
	logrusKit.DisableQuote(nil)
}

// @title Title
// @version 1.0
// @description Description.
// @query.collection.format multi
func main() {
	jsonRespKit.MustSetUp(func(code, msg string, data interface{}) interface{} {
		return &types.JsonResponse{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	})

	pool, err := poolKit.NewPool(2000)
	if err != nil {
		logrus.Fatal(err)
	}
	pushKit.MustSetUp(pool, nil, time.Second*5)

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

	// close
	{
		engine.POST("/close_by_id", ginKit.WrapToHandlerFunc(handler.CloseById))
		engine.POST("/close_by_bsid", ginKit.WrapToHandlerFunc(handler.CloseByBsid))
		engine.POST("/close_all", ginKit.WrapToHandlerFunc(handler.CloseAll))
		engine.POST("/close_by_user", ginKit.WrapToHandlerFunc(handler.CloseByUser))
		engine.POST("/close_by_group", ginKit.WrapToHandlerFunc(handler.CloseByGroup))
	}

	// WebSocket
	{
		processor, err := wsKit.NewProcessor(nil, nil, &types.Listener{}, wsKit.MessageTypeText)
		if err != nil {
			logrus.Fatal(err)
		}
		engine.GET("/ws", processor.ProcessWithGin)
	}

	// SSE
	{
		processor, err := sseKit.NewProcessor(nil, &types.Listener{}, sseKit.MessageTypeRaw)
		if err != nil {
			logrus.Fatal(err)
		}
		engine.GET("/sse", processor.ProcessWithGin)
	}

	// html
	{
		// 传参路径root 是相对于项目的根目录(working directory)，而非main()所在的目录（虽然他们常常是同一个）
		// "./internal/pushDemo/web" <=> "internal/pushDemo/web"
		root := "internal/pushDemo/web"
		if err := ginKit.StaticDir(engine, "/s", root, false); err != nil {
			logrus.Fatal(err)
		}
	}

	ginKit.DefaultFavicon(engine)
	if err := ginKit.DefaultNoRoute(engine); err != nil {
		logrus.Fatal(err)
	}

	// 每隔10s输出 statistics
	go func() {
		c := cronKit.NewCron()
		_, err := c.AddFunc("@every 10s", func() {
			logrus.Infof("statistics:\n%s\n", pushKit.GetStatistics())
		})
		if err != nil {
			logrus.Fatal(err)
		}
		c.Run()
	}()

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}
