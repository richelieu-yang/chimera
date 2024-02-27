package i18nPackKit

import (
	"github.com/duke-git/lancet/v2/condition"
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
	innerMaker  Maker

	defaultMaker = func(code, msg string, data interface{}) interface{} {
		return &bean{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	}
)

func MustSetUp(defaultLanguage language.Tag, messageFilePaths []string, maker Maker) {
	if err := SetUp(defaultLanguage, messageFilePaths, maker); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(defaultLanguage language.Tag, messageFilePaths []string, maker Maker) (err error) {
	defer func() {
		if err != nil {
			innerBundle = nil
			innerMaker = nil
		}
	}()

	innerBundle, err = i18nKit.NewBundle(defaultLanguage, messageFilePaths...)
	innerMaker = maker
	return
}

func getMaker() Maker {
	return condition.TernaryOperator(innerMaker != nil, innerMaker, defaultMaker)
}
