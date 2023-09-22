package sessionKit

import (
	"github.com/gorilla/sessions"
	"net/http"
)

// GetSession 尝试获取一个 *sessions.Session 实例（优先用已有的，没的话再创建）.
/*
Description: 直接使用 源 即可.

PS: 可以通过 sessions.Session 的 IsNew 属性判断该实例是"新创建的"还是"原有的".

@param req 			e.g. gin中的ctx.Request
@param cookieName	浏览器端cookie的name
*/
func GetSession(req *http.Request, store sessions.Store, cookieName string) (*sessions.Session, error) {
	return store.Get(req, cookieName)
}

// SaveSession
/*
Description: 直接使用 源 即可.

@param req		e.g. gin中的ctx.Request
@param writer 	e.g. gin中的ctx.Writer
*/
func SaveSession(req *http.Request, writer http.ResponseWriter, session *sessions.Session) error {
	return session.Save(req, writer)
}
