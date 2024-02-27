package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"golang.org/x/text/language"
)

// Associate 关联
/*
@param languageCode 语言代码，可参考: i18nKit/_info.md
*/
func Associate(bundle *i18n.Bundle, languageCode string, messageFile *i18n.MessageFile) error {
	if err := interfaceKit.AssertNotNil(bundle, "bundle"); err != nil {
		return err
	}
	if err := strKit.AssertNotEmpty(languageCode, "languageCode"); err != nil {
		return err
	}
	if err := interfaceKit.AssertNotNil(messageFile, "messageFile"); err != nil {
		return err
	}

	tag := language.Make(languageCode)
	return bundle.AddMessages(tag, messageFile.Messages...)
}
