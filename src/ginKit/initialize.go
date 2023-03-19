package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/mainControl"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// Initialize 初始化gin的配置.
func Initialize(colorful bool) {
	// 设置模式（gin实际上有3种模式）
	if mainControl.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 日志颜色
	if colorful {
		// 强制设置日志颜色
		gin.ForceConsoleColor()
	} else {
		// 禁止日志颜色
		gin.DisableConsoleColor()
	}

	// 通过logrus输出Gin的日志
	// Richelieu：从目前表现来看，虽然gin和logrus都可以设置颜色，但在此处,只要gin允许了，logrus的logger是否允许就无效了
	logger = logrusKit.NewCustomizedLogger(nil, mainControl.GetLogrusLevel())
	gin.DefaultWriter = logger.Out
}

// OnDebugChanged debug发生了改变，进行相应处理
func OnDebugChanged() {
	if mainControl.IsDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	if logger != nil {
		logger.SetLevel(mainControl.GetLogrusLevel())
		gin.DefaultWriter = logger.Out
	}
}
