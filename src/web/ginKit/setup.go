package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/netKit"
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config, recoveryMiddleware gin.HandlerFunc, businessLogic func(engine *gin.Engine) error) {
	err := setUp(config, recoveryMiddleware, businessLogic)
	if err != nil {
		logrus.Fatal(err)
	}
}

// setUp
/*
PS: 正常执行的情况下，此方法会阻塞调用的协程.

@param config				可以为nil（将返回error）
@param recoveryMiddleware 	可以为nil（将采用默认值 gin.Recovery()）
@param businessLogic 		可以为nil；业务逻辑，可以在其中进行 路由绑定 等操作...
*/
func setUp(config *Config, recoveryMiddleware gin.HandlerFunc, businessLogic func(engine *gin.Engine) error) error {
	if err := config.CheckAndPolyfill(); err != nil {
		return err
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
	// 通过logrus输出Gin的日志.
	// Richelieu：从目前表现来看，虽然gin和logrus都可以设置颜色，但在此处,只要gin允许了，logrus的logger是否允许就无效了
	gin.DefaultWriter = logrus.StandardLogger().Out

	engine := NewEngine()
	/*
		MaxMultipartMemory只是限制内存，不是针对文件上传文件大小，即使文件大小比这个大，也会写入临时文件。
		默认32MiB，并不涉及"限制上传文件的大小"，原因：上传的文件s按顺序存入内存中，累加大小不得超出 32Mb ，最后累加超出的文件就存入系统的临时文件中。非文件字段部分不计入累加。所以这种情况，文件上传是没有任何限制的。
		参考: https://studygolang.com/articles/22643
	*/
	engine.MaxMultipartMemory = 32 << 20

	// middleware
	if err := attachMiddlewares(engine, config.Middleware, recoveryMiddleware); err != nil {
		return err
	}

	if businessLogic != nil {
		if err := businessLogic(engine); err != nil {
			return err
		}
	}

	// http server
	go func() {
		if err := engine.Run(netKit.JoinHostnameAndPort(config.Host, config.Port)); err != nil {
			logrus.Fatal(engine)
		}
	}()

	// https server
	sslConfig := config.SSL
	if sslConfig != nil {
		go func() {
			if err := engine.RunTLS(netKit.JoinHostnameAndPort(config.Host, sslConfig.Port), sslConfig.CertFile, sslConfig.KeyFile); err != nil {
				logrus.Fatal(engine)
			}
		}()
	}

	select {}
}
