package componentKit

import (
	"gitee.com/richelieu042/go-scales/src/jsonKit"
	"gitee.com/richelieu042/go-scales/src/msgKit"
	"github.com/sirupsen/logrus"
)

// InitializeJsonComponent 初始化json组件（可选）
/*
@param msgProcessor		回调函数（可用于修改响应给前端实例中的msg属性）
@param messageFilePath	（存储code和msg映射关系的）文件的路径（相对 || 绝对）
*/
func InitializeJsonComponent(msgProcessor jsonKit.MsgProcessor, messageFilePath string) error {
	jsonKit.SetMsgProcessor(msgProcessor)
	err := msgKit.ReadFile(messageFilePath)
	if err == nil {
		logrus.Info("[COMPONENT, JSON] Initialize successfully.")
	}
	return err
}
