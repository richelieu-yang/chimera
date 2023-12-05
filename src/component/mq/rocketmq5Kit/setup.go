package rocketmq5Kit

import (
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
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
@param verifyTopic		可以为""（不进行验证）
*/
func SetUp(c *Config, clientLogPath string, verifyConfig *VerifyConfig) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	if err = check(c); err != nil {
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

func check(c *Config) error {
	if c == nil {
		return errorKit.New("c == nil")
	}

	// Endpoints
	s := sliceKit.RemoveEmpty(c.Endpoints, true)
	if err := sliceKit.AssertNotEmpty(s, "c.Endpoints"); err != nil {
		return err
	}
	c.Endpoints = s

	// Credentials
	if c.Credentials == nil {
		c.Credentials = &credentials.SessionCredentials{
			AccessKey:     "",
			AccessSecret:  "",
			SecurityToken: "",
		}
	}

	return nil
}
