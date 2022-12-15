package componentKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/ginKit"
	"github.com/richelieu42/go-scales/src/ipKit"
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/richelieu42/go-scales/src/mainControl"
	"github.com/richelieu42/go-scales/src/netKit"
)

// InitializeGinComponent 初始化Gin组件（可选）
/*
PS: 由于会阻塞当前goroutine，建议在业务的最后调用此方法.

@param recoveryMiddleware 	可以为nil，将采用默认值 gin.Recovery()
@param businessLogic 		可以为nil；用于添加业务逻辑，可以在此回调函数内：设置Gin的模式（gin.SetMode()）、对"中间件"和"路由"进行定制化处理...
*/
func InitializeGinComponent(recoveryMiddleware gin.HandlerFunc, businessLogic func(engine *gin.Engine) error) error {
	config, err := GetGinConfig()
	if err != nil {
		return err
	}
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	/* 配置检查 */
	switch config.Host {
	case "":
		fallthrough
	case "localhost":
	default:
		if !ipKit.IsIPv4(config.Host) {
			return errorKit.Simple("host(%s) of gin is invalid", config.Host)
		}
	}
	_, err = netKit.IsLocalPortUsable(config.Port)
	if err != nil {
		return err
	}

	/* Gin的模式，默认debug模式，后续可以在 businessLogic 里面调整 */
	gin.SetMode(gin.DebugMode)

	/* 日志颜色 */
	if config.Colorful {
		// 强制设置日志颜色
		gin.ForceConsoleColor()
	} else {
		// 禁止日志颜色
		gin.DisableConsoleColor()
	}

	/* 通过logrus输出Gin的日志 */
	// Richelieu：从目前表现来看，虽然gin和logrus都可以设置颜色，但在此处,只要gin允许了，logrus的logger是否允许就无效了
	logger := logrusKit.NewCustomizedLogger(nil, mainControl.GetLogrusLevel())
	gin.DefaultWriter = logger.Out

	engine := ginKit.NewEngine()

	/* middleware */
	if err := ginKit.AttachCommonMiddlewares(engine, config.Middleware, recoveryMiddleware); err != nil {
		return err
	}

	/* businessLogic */
	if businessLogic != nil {
		if err := businessLogic(engine); err != nil {
			return err
		}
	}

	/*
		启动服务:
		会阻塞当前协程（一般是main程）；
		Richelieu: 此处也可以选择通过 goroutine 来启动gin服务，但个人感觉没这个必要.
	*/
	return engine.Run(netKit.JoinHostPort(config.Host, config.Port))
}
