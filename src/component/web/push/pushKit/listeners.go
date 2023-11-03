package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
)

var inner = &types.InnerListener{}

// NewListeners
/*
PS: 本方法仅供本项目使用，严禁外部调用.
*/
func NewListeners(listener types.Listener) (types.Listeners, error) {
	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return []types.Listener{inner, listener}, nil
}
