package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func UnBindId(channel Channel) {
	id := channel.GetId()

	/* 写锁 */
	idMap.LockFunc(func() {
		delete(idMap.Map, id)
	})
}

func UnbindBsid(channel Channel) {
	bsid := channel.GetBsid()
	if strKit.IsEmpty(bsid) {
		return
	}

	defer channel.ClearBsid()

	/* 写锁 */
	bsidMap.LockFunc(func() {
		delete(bsidMap.Map, bsid)
	})
}

// UnbindUser
/*
PS: 解绑成功后，如果set为空，应该移除掉.
*/
func UnbindUser(channel Channel) {
	user := channel.GetUser()
	if strKit.IsEmpty(user) {
		return
	}
	defer channel.ClearUser()

	userSet := GetUserSet(user)
	if userSet == nil {
		return
	}

	/* 写锁 */
	userSet.LockFunc(func() {
		userSet.Set.Remove(channel)
	})
}

// UnbindGroup
/*
PS: 解绑成功后，如果set为空，应该移除掉.
*/
func UnbindGroup(channel Channel) {
	group := channel.GetGroup()
	if strKit.IsEmpty(group) {
		return
	}
	defer channel.ClearGroup()

	groupSet := GetGroupSet(group)
	if groupSet == nil {
		return
	}

	/* 写锁 */
	groupSet.LockFunc(func() {
		groupSet.Set.Remove(channel)
	})
}
