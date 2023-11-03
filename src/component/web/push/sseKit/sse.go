package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
)

// NewProcessor
/*
@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
*/
func NewProcessor(idGenerator func() (string, error), listener types.Listener, msgType messageType) (pushKit.Processor, error) {
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}
	listeners, err := types.NewListeners(listener)
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
