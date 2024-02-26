package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/language"
)

var (
	NotSetupError = errorKit.New("Haven’t been set up correctly")
)

var (
	innerBundle *i18n.Bundle
)

func MustSetUp(defaultLanguage language.Tag, messageFilePaths ...string) {
	if err := SetUp(defaultLanguage, messageFilePaths...); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(defaultLanguage language.Tag, messageFilePaths ...string) error {
	bundle, err := NewBundle(defaultLanguage, messageFilePaths...)
	if err != nil {
		return err
	}

	innerBundle = bundle
	return nil
}

// GetMessage
/*
@param langs 可以不传，此时将使用 innerBundle 的默认语言
*/
func GetMessage(messageId string, langs ...string) (string, error) {
	if innerBundle == nil {
		return "", NotSetupError
	}

	return GetMessageByBundle(innerBundle, messageId, langs...)
}
