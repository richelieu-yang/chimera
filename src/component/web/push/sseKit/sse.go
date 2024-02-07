package sseKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
)

// NewProcessor
/*
@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
*/
func NewProcessor(idGenerator func() (string, error), listener pushKit.Listener, msgType messageType) (pushKit.Processor, error) {
	if idGenerator == nil {
		idGenerator = pushKit.DefaultIdGenerator()
	}
	listeners, err := pushKit.NewListeners(listener, true)
	if err != nil {
		return nil, err
	}

	processor := &SseProcessor{
		idGenerator: idGenerator,
		listeners:   listeners,
		msgType:     msgType,
	}
	return processor, nil
}
