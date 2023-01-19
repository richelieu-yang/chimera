package sessionKit

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
	"testing"
)

// TestCookieStore
/*
访问地址: http://localhost/test
*/
func TestCookieStore(t *testing.T) {
	cookieName := "session-id"

	keyPairs := []byte("123456789")
	options := &sessions.Options{
		Path:   "/",
		MaxAge: 1800,
	}

	store := NewCookieStore(keyPairs, options)

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		/* (1) 获取session */
		session, err := store.Get(ctx.Request, cookieName)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		/* (2) 保存session数据，本质上是将内存中的数据持久化到存储介质中（序列化并写到Redis中；会重置key的TTL） */
		if err := session.Save(ctx.Request, ctx.Writer); err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ctx.String(http.StatusOK, "set IsNew: [%t].", session.IsNew)
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
