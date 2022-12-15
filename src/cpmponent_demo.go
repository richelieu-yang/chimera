package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/component/componentKit"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"net/http"
	"time"
)

func main() {
	if err := componentKit.InitializeEnvironment(); err != nil {
		errorKit.PanicByError(err)
	}

	// 业务逻辑
	if err := business(); err != nil {
		errorKit.PanicByError(err)
	}

	//// redis组件（可选）
	//if err := componentKit.InitializeRedisComponent(); err != nil {
	//	errorKit.PanicByError(err)
	//}

	// json组件（可选）
	var msgProcessor jsonKit.MessageProcessor = func(code string, msg string, data interface{}) string {
		// TODO: 对响应结构体中的message进行二开，比如可以加上: 是哪台服务响应的
		return msg
	}
	var respProcessor jsonKit.ResponseProcessor = func(resp *jsonKit.Response) any {
		// TODO: 对响应结构体进行二开，修改序列化为json字符串时的key
		return resp
	}
	if err := componentKit.InitializeJsonComponent("", msgProcessor, respProcessor); err != nil {
		errorKit.PanicByError(err)
	}

	// gin组件（可选）
	recoveryMiddleware := gin.CustomRecovery(func(c *gin.Context, err any) {
		// TODO: gin处理请求时发生panic的情况，在此处进行相应的处理（比如响应json给前端）
	})
	if err := componentKit.InitializeGinComponent(recoveryMiddleware, routeBusiness); err != nil {
		errorKit.PanicByError(err)
	}
}

// 业务逻辑1: 读取业务配置文件...
func business() error {

	return nil
}

// 业务逻辑2: 绑定路由...
func routeBusiness(engine *gin.Engine) error {
	// test
	engine.Any("/test.act", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, time.Now().UTC().Format(time.RFC3339))
	})

	return nil
}
