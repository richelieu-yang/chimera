package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func UnBindId(channel Channel) {
	id := channel.GetId()

	/* 写锁 */
	idMap.RWLock.LockFunc(func() {
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
	bsidMap.RWLock.LockFunc(func() {
		delete(bsidMap.Map, bsid)
	})
}

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
	userSet.RWLock.LockFunc(func() {
		userSet.Set.Remove(channel)
	})
}

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
	groupSet.RWLock.LockFunc(func() {
		groupSet.Set.Remove(channel)
	})
}
