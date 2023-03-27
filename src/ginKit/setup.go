package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/netKit"
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config, recoveryMiddleware gin.HandlerFunc, businessLogic func(engine *gin.Engine) error) {
	err := SetUp(config, recoveryMiddleware, businessLogic)
	if err != nil {
		logrus.Fatal(err)
	}
}

// SetUp
/*
PS: 正常执行的情况下，此方法会阻塞调用的协程.

@param recoveryMiddleware 	可以为nil（将采用默认值 gin.Recovery()）
@param businessLogic 		可以为nil；业务逻辑，可以在其中进行 路由绑定 等操作...
*/
func SetUp(config *Config, recoveryMiddleware gin.HandlerFunc, businessLogic func(engine *gin.Engine) error) error {
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	// Gin的模式，默认debug模式，后续可以在 businessLogic 里面调整
	gin.SetMode(gin.DebugMode)
	// 日志颜色
	if config.Colorful {
		// 强制设置日志颜色
		gin.ForceConsoleColor()
	} else {
		// 禁止日志颜色
		gin.DisableConsoleColor()
	}
	// 通过logrus输出Gin的日志. Richelieu：从目前表现来看，虽然gin和logrus都可以设置颜色，但在此处,只要gin允许了，logrus的logger是否允许就无效了
	logger := logrusKit.NewLogger(nil, logrus.DebugLevel)
	gin.DefaultWriter = logger.Out

	engine := NewEngine()
	// middleware
	if err := AttachCommonMiddlewares(engine, config.Middleware, recoveryMiddleware); err != nil {
		return err
	}
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
	return engine.Run(netKit.JoinHostnameAndPort(config.Host, config.Port))
}
