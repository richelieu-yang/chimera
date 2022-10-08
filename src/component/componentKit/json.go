package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"github.com/richelieu42/go-scales/src/msgKit"
	"github.com/sirupsen/logrus"
)

// InitializeJsonComponent 初始化json组件（可选）
/*
@param msgProcessor		回调函数（可用于修改响应给前端实例中的msg属性）
@param messageFilePath	（存储code和msg映射关系的）文件的路径（相对 || 绝对），如果为空则不读取message文件
*/
func InitializeJsonComponent(msgProcessor jsonKit.MsgProcessor, messageFilePath string) error {
	jsonKit.SetMsgProcessor(msgProcessor)

	if strKit.IsEmpty(messageFilePath) {
		logrus.Warn("[COMPONENT, JSON] messageFilePath is empty.")
	} else {
		if err := msgKit.ReadFile(messageFilePath); err != nil {
			return err
		}
	}

	logrus.Info("[COMPONENT, JSON] Initialize successfully.")
	return nil
}
