package i18nPackKit

import (
	"github.com/duke-git/lancet/v2/condition"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/i18n/i18nKit"
	"golang.org/x/text/language"
)

var (
	NotSetupError = errorKit.Newf("Haven’t been set up correctly")
)

var (
	innerBundle    *i18n.Bundle
	innerBeanMaker BeanMaker

	defaultBeanMaker = func(code, msg string, data interface{}) interface{} {
		return &bean{
			Code:    code,
			Message: msg,
			Data:    data,
		}
	}
)

func SetUp(defaultLanguage language.Tag, maker BeanMaker) {
	innerBundle = i18nKit.NewBundle(defaultLanguage)
	innerBeanMaker = maker
}

func getMaker() BeanMaker {
	return condition.TernaryOperator(innerBeanMaker != nil, innerBeanMaker, defaultBeanMaker)
}
