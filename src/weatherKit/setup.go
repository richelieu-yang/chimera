package weatherKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var apiKey string

func MustSetUp(key string) {
	err := setUp(key)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// setUp
/*
PS: 正常执行的情况下，此方法会阻塞调用的协程.

@param config				可以为nil（将返回error）
@param recoveryMiddleware 	可以为nil（将采用默认值 gin.Recovery()）
@param businessLogic 		可以为nil；业务逻辑，可以在其中进行 路由绑定 等操作...
*/
func setUp(key string) error {
	if err := strKit.AssertNotBlank(key, "key"); err != nil {
		return err
	}

	apiKey = key
	return nil
}
