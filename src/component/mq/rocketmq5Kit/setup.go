package rocketmq5Kit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
	"github.com/sirupsen/logrus"
)

var (
	NotSetupError = errorKit.New("Haven’t been set up correctly")
)

var config *Config

func MustSetUp(c *Config, clientLogPath string, verifyConfig *VerifyConfig) {
	if err := SetUp(c, clientLogPath, verifyConfig); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// SetUp
/*
@param clientLogPath	可以为""（输出到控制台）
@param verifyConfig		可以为nil（不进行验证）
*/
func SetUp(c *Config, clientLogPath string, verifyConfig *VerifyConfig) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	if err = validateKit.Struct(c); err != nil {
		return
	}

	// 客户端日志输出
	if err = setClientLog(clientLogPath); err != nil {
		return
	}

	config = c

	// verify
	if err = verify(verifyConfig); err != nil {
		return
	}

	return
}
