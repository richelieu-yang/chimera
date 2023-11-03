package pushKit

import "github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"

func UnBindId(channel types.Channel) {
	id := channel.GetId()

	// 写锁
	idMap.RWLock.LockFunc(func() {
		delete(idMap.Map, id)
	})
}

func UnbindBsid(channel types.Channel) {
	bsid := channel.GetBsid()

	// 写锁
	bsidMap.RWLock.LockFunc(func() {
		delete(bsidMap.Map, bsid)
	})
}

func UnbindUser(channel types.Channel) {
	user := channel.GetUser()

	userSet := GetUserSet(user)
	if userSet == nil {
		return
	}

	// 写锁
	userSet.RWLock.LockFunc(func() {
		userSet.Set.Remove(channel)
	})
}

func UnbindGroup(channel types.Channel) {
	group := channel.GetGroup()

	groupSet := GetGroupSet(group)
	if groupSet == nil {
		return
	}

	// 写锁
	groupSet.RWLock.LockFunc(func() {
		groupSet.Set.Remove(channel)
	})
}
