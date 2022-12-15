package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"github.com/richelieu42/go-scales/src/msgKit"
	"github.com/sirupsen/logrus"
)

// InitializeJsonComponent 初始化json组件（可选）
/*
@param messageFilePath	（存储code和msg映射关系的）文件的路径（相对 || 绝对），如果为空则不读取message文件
@param msgProcessor	可以为nil，对响应结构体中的message进行二开，比如可以加上: 是哪台服务响应的
@param respProcess	可以为nil，对响应结构体进行二开，以修改序列化为json字符串时的key
*/
func InitializeJsonComponent(messageFilePath string, messageProcessor jsonKit.MessageProcessor, respProcessor jsonKit.ResponseProcessor) error {
	if strKit.IsEmpty(messageFilePath) {
		logrus.Warn("[COMPONENT, JSON] messageFilePath is empty.")
	} else {
		if err := msgKit.ReadFile(messageFilePath); err != nil {
			return err
		}
	}

	jsonKit.SetMsgProcessor(messageProcessor)
	jsonKit.SetRespProcessor(respProcessor)

	logrus.Info("[COMPONENT, JSON] Initialize successfully.")
	return nil
}
