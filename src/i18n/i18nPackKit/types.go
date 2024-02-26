package i18nPackKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

type (
	Packager struct {
		langs []string
	}
)

func NewPackager(langs ...string) *Packager {
	return &Packager{
		langs: langs,
	}
}

func (p *Packager) Pack(code string, data interface{}, msgArgs ...interface{}) interface{} {
	return PackFully(code, "", data, msgArgs...)
}

func (p *Packager) PackFully(code, msg string, data interface{}, msgArgs ...interface{}) interface{} {
	if strKit.IsEmpty(msg) {
		msg = msgMap[code]
	}
	if strKit.IsNotEmpty(msg) && msgArgs != nil {
		msg = fmt.Sprintf(msg, msgArgs...)
	}
	return provider(code, msg, data)
}
