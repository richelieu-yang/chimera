package pushKit

func UnBindId(channel Channel, id string) {
	// 写锁
	idMap.RWLock.LockFunc(func() {
		delete(idMap.Map, id)
	})
}

func UnbindBsid(channel Channel, bsid string) {
	// 写锁
	bsidMap.RWLock.LockFunc(func() {
		delete(bsidMap.Map, bsid)
	})
}

func UnbindUser(channel Channel, user string) {
	userSet := GetUserSet(user)
	if userSet == nil {
		return
	}

	// 写锁
	userSet.RWLock.LockFunc(func() {
		userSet.Set.Remove(channel)
	})
}

func UnbindGroup(channel Channel, group string) {
	groupSet := GetGroupSet(group)
	if groupSet == nil {
		return
	}

	// 写锁
	groupSet.RWLock.LockFunc(func() {
		groupSet.Set.Remove(channel)
	})
}
