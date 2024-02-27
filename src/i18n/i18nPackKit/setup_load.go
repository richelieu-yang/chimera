package i18nPackKit

import "github.com/richelieu-yang/chimera/v3/src/i18n/i18nKit"

func LoadMessageFile(path string) error {
	if innerBundle == nil {
		return NotSetupError
	}

	return i18nKit.LoadMessageFile(innerBundle, path)
}

func LoadMessageFileBytes(content []byte, path string) error {
	if innerBundle == nil {
		return NotSetupError
	}

	return i18nKit.LoadMessageFileBytes(innerBundle, content, path)
}
