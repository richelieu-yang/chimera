package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func CloseById(id string, reason string) (err error) {
	reason = strKit.EmptyToDefault(reason, "no reason")

	var channel Channel
	idMap.LockFunc(func() {
		channel = idMap.Map[id]
		if channel == nil {
			err = errorKit.New("No channel for id(%s)", id)
			return
		}
		err = channel.Close(reason)
		// 解绑（unbind）后续由 inner handler 处理
	})
	return
}
