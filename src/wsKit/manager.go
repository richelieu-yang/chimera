package wsKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// key: group
var groupMapper = make(map[string][]*Connection)

// key: user
var userMapper = make(map[string][]*Connection)

// key: uniqueId
var all = make(map[string][]*Connection)

//var rwLock = new(sync.RWMutex)

func addToMapper(mapper map[string][]*Connection, key string, conn *Connection) {
	s := mapper[key]
	if s == nil {
		s = make([]*Connection, 16)
	}

	mapper[key] = s
}

func removeFromMapper(mapper map[string][]*Connection) {

}

// BindGroup 可以多次绑定
func BindGroup(conn *Connection, group string) {
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

func UnbindGroup(conn *Connection) {
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

func BindUser(conn *Connection, user string) {
	if conn == nil {
		return
	}

	conn.user = user
}

func BindUniqueId(conn *Connection, uniqueId string) {
	if conn == nil {
		return
	}

	conn.uniqueId = uniqueId
}
