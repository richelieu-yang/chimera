package gaodeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var apiKey string
var NotSetupError = errorKit.New("Haven’t been set up correctly")

func MustSetUp(key string) {
	err := setUp(key)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func setUp(key string) error {
	if err := strKit.AssertNotBlank(key, "key"); err != nil {
		return err
	}

	apiKey = key
	return nil
}

func GetApiKey() (string, error) {
	if strKit.IsEmpty(apiKey) {
		return "", NotSetupError
	}
	return apiKey, nil
}
