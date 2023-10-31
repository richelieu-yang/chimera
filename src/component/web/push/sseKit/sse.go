package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
)

// NewProcessor
/*
@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
*/
func NewProcessor(idGenerator func() (string, error), listener pushKit.Listener, msgType messageType) (pushKit.Processor, error) {
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}
	listeners, err := pushKit.NewListeners(listener)
	if err != nil {
		return nil, err
	}
	switch msgType {
	case MessageTypeRaw:
	case MessageTypeEncode:
	case MessageTypeBase64:
	default:
		return nil, errorKit.New("invalid message type(%d)", msgType)
	}

	processor := &SseProcessor{
		idGenerator: idGenerator,
		listeners:   listeners,
		msgType:     msgType,
	}
	return processor, nil
}
