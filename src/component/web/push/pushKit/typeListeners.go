package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
)

// NewListeners
/*
PS: 本方法仅供本项目使用，严禁外部调用.
*/
func NewListeners(listener Listener, sseFlag bool) (Listeners, error) {
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	inner := &innerListener{
		sseFlag: sseFlag,
	}
	return []Listener{inner, listener}, nil
}
