package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
)

func NewProcessor(idGenerator func() (string, error), listener pushKit.Listener, messageType messageType) (pushKit.Processor, error) {

}
