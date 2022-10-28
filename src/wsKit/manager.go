package wsKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// key: group
var groupMapper = make(map[string][]*Channel)

// key: user
var userMapper = make(map[string][]*Channel)

// key: uniqueId
var all = make(map[string][]*Channel)

//var rwLock = new(sync.RWMutex)

func addToMapper(mapper map[string][]*Channel, key string, conn *Channel) {
	s := mapper[key]
	if s == nil {
		s = make([]*Channel, 16)
	}

	mapper[key] = s
}

func removeFromMapper(mapper map[string][]*Channel) {

}

// BindGroup 可以多次绑定
func BindGroup(conn *Channel, group string) {
	if conn == nil {
		return
	}
	group = strKit.Trim(group)
	if strKit.IsEmpty(group) {
		// 直接解绑
		UnbindGroup(conn)
		return
	}

	if group == conn.group {
		// 前后的group一致
		return
	}
	delete(groupMapper, conn.group)
	conn.group = group
	//groupMapper[conn.group] = conn

	//oldGroup := conn.group
	//if strKit.IsNotEmpty(oldGroup) {
	//
	//}
	//
	//conn.group = group
}

func UnbindGroup(conn *Channel) {
	if conn == nil {
		return
	}

	old := conn.group
	if strKit.IsEmpty(old) {
		// 本来就没绑定，无需解绑
		return
	}
	// 解绑
	delete(groupMapper, old)
}

func BindUser(conn *Channel, user string) {
	if conn == nil {
		return
	}

	conn.user = user
}

func BindUniqueId(conn *Channel, uniqueId string) {
	if conn == nil {
		return
	}

	conn.uniqueId = uniqueId
}
