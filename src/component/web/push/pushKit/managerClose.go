package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// CloseById
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseById(id string, reason string) (err error) {
	reason = strKit.EmptyToDefault(reason, "no reason")

	/* 写锁 */
	idMap.LockFunc(func() {
		channel := idMap.Map[id]
		if channel == nil {
			return
		}
		err = channel.Close(reason)
	})
	return
}

// CloseByBsid
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseByBsid(bsid string, reason string) (err error) {
	reason = strKit.EmptyToDefault(reason, "no reason")

	/* 写锁 */
	bsidMap.LockFunc(func() {
		channel := bsidMap.Map[bsid]
		if channel == nil {
			return
		}
		err = channel.Close(reason)
		// 解绑（unbind）后续由 inner handler 处理
	})
	return
}

// CloseByUser
/*
PS: 解绑（unbind）后续由 inner handler 处理.
*/
func CloseByUser(user string, reason string) {
	reason = strKit.EmptyToDefault(reason, "no reason")

	/* 写锁 */
	userMap.LockFunc(func() {
		userSet := userMap.Map[user]

		/* 写锁 */
		userSet.LockFunc(func() {
			userSet.Set.Each(func(channel Channel) bool {
				_ = channel.Close(reason)
				return false
			})
		})
	})
}
