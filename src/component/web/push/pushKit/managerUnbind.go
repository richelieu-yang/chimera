package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// UnbindId
/*
PS: 仅是解绑，不会关闭channel.
*/
func UnbindId(channel Channel) {
	id := channel.GetId()

	/* 写锁 */
	idMap.LockFunc(func() {
		delete(idMap.Map, id)
	})
}

// UnbindBsid
/*
PS: 仅是解绑，不会关闭channel.
*/
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
PS:
(1) 仅是解绑，不会关闭channel;
(2) 解绑成功后，如果set为空，应该移除掉.
*/
func UnbindUser(channel Channel) {
	user := channel.GetUser()
	if strKit.IsEmpty(user) {
		return
	}
	defer channel.ClearUser()

	/* map写锁 */
	userMap.LockFunc(func() {
		userSet, ok := userMap.Map[user]
		if !ok {
			return
		}
		if userSet == nil {
			delete(userMap.Map, user)
			return
		}
		/* set写锁 */
		userSet.LockFunc(func() {
			userSet.Set.Remove(channel)
			if userSet.Set.Cardinality() == 0 {
				delete(userMap.Map, user)
			}
		})
	})
}

// UnbindGroup
/*
PS:
(1) 仅是解绑，不会关闭channel;
(2) 解绑成功后，如果set为空，应该移除掉.
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
