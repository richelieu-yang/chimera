package pushKit

func UnbindGroup(channel Channel, group string) {

}

func UnbindUser(channel Channel, user string) {

}

func UnbindBsid(channel Channel, bsid string) {

}

func UnBindId(channel Channel, id string) {
	// 写锁
	allMap.RWLock.LockFunc(func() {
		delete(allMap.Map, id)
	})
}
