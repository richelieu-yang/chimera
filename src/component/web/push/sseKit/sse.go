package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	types2 "github.com/richelieu-yang/chimera/v2/src/component/web/push/sseKit/types"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
)

// NewProcessor
/*
@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
*/
func NewProcessor(idGenerator func() (string, error), listener types.Listener, msgType types2.MessageType) (types.Processor, error) {
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}
	listeners, err := types.NewListeners(listener)
	if err != nil {
		return nil, err
	}

	processor := &types2.SseProcessor{
		idGenerator: idGenerator,
		listeners:   listeners,
		msgType:     msgType,
	}
	return processor, nil
}
