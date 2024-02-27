package i18nPackKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/i18n/i18nKit"
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

func MustSetUp(defaultLanguage language.Tag, messageFilePaths []string) {
	if err := SetUp(defaultLanguage, messageFilePaths); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(defaultLanguage language.Tag, messageFilePaths []string) (err error) {
	defer func() {
		if err != nil {
			innerBundle = nil
		}
	}()

	innerBundle, err = i18nKit.NewBundle(defaultLanguage, messageFilePaths...)
	return
}
