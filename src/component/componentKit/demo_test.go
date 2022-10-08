package componentKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"testing"
)

func Test(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	if err := InitializeEnvironment(); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "env", err)
	}

	// redis组件（可选）
	if err := InitializeRedisComponent(); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "redis", err)
	}

	if err := tmp(); err != nil {
		errorKit.PanicByError(err)
	}

	msgProcessor := func(code string, msg string, data interface{}) string {
		return strKit.Format("[%s] %s", code, msg)
	}
	if err := InitializeJsonComponent(msgProcessor, ""); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "json", err)
	}

	// gin组件（可选）
	if err := InitializeGinComponent(nil, tmp1); err != nil {
		errorKit.Panic("fail to initialize %s, error: %+v", "gin", err)
	}
}

func tmp() error {
	// 业务逻辑（读取配置文件...）
	return nil
}

func tmp1(engine *gin.Engine) error {
	// 业务逻辑（绑定路由...）
	return nil
}
