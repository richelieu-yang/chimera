package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
)

var inner = &InnerListener{}

// NewListeners
/*
PS: 本方法仅供本项目使用，严禁外部调用.
*/
func NewListeners(listener Listener) (Listeners, error) {
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return []Listener{inner, listener}, nil
}
