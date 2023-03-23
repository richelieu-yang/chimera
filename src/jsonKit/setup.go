package jsonKit

import (
	"github.com/richelieu42/chimera/src/msgKit"
	"github.com/sirupsen/logrus"
	"sync"
)

var setupOnce sync.Once

func MustSetUp(messageHook MessageHook, responseHook ResponseHook, messageFiles ...string) {
	err := SetUp(messageHook, responseHook, messageFiles...)
	if err != nil {
		logrus.Fatal(err)
	}
}

// SetUp
/*
@param msgProcessor		[可以为nil] 对响应结构体中的message进行二开，比如可以加上: 是哪台服务响应的
@param respProcess		[可以为nil] 对响应结构体进行二开，以修改序列化为json字符串时的key
@param messageFiles		[.properties文件] （存储code和msg映射关系的）文件的路径（相对 || 绝对），如果为空则不读取message文件
*/
func SetUp(msgHook MessageHook, respHook ResponseHook, messageFiles ...string) (err error) {
	setupOnce.Do(func() {
		for _, path := range messageFiles {
			err = msgKit.ReadFile(path)
			if err != nil {
				return
			}
		}

		SetMsgHook(msgHook)
		SetRespHook(respHook)
	})

	return
}
