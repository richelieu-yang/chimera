package pushKit

import "github.com/richelieu-yang/chimera/v2/src/core/strKit"

func CloseById(id string, reason string) (err error) {
	reason = strKit.EmptyToDefault(reason, "no reason")

	idMap.LockFunc(func() {
		channel := idMap.Map[id]
		if channel != nil {
			err = channel.Close(reason)
			// 解绑后续由 inner handler 处理
		}
	})
	return
}
