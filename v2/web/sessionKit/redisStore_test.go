package sessionKit

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/richelieu42/chimera/v2/idKit"
	"net/http"
	"testing"
	"time"
)

// TestRedisStore
/*
准备工作: 启动Redis.
访问地址: http://localhost/test
*/
func TestRedisStore(t *testing.T) {
	// Redis中的key的前缀（value为 string 类型）
	redisKeyPrefix := "session:"

	cookieName := "session-id"

	// Redis配置（单节点）
	redisOptions := &redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	}

	client := redis.NewClient(redisOptions)

	store, err := NewRedisStore(context.TODO(), client)
	if err != nil {
		panic(err)
	}
	//defer store.Close()
	// 自定义: Redis中的key的前缀
	store.KeyPrefix(redisKeyPrefix)
	// 自定义: cookie的配置
	store.Options(sessions.Options{
		HttpOnly: true,
		Secure:   false,
		MaxAge:   0, // 只有 > 0 的情况下，才会将数据写到Redis中
	})
	// 自定义: cookie的value、Redis中的key的后半部分
	store.KeyGen(func() (string, error) {
		return idKit.NewULID(), nil
	})
	store.SetSessionTimeoutWhenMaxAgeZero(time.Hour)

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
