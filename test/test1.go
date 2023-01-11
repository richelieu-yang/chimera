package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/idKit"
	"net/http"
)

var (
	// Redis中的key的前缀（value为 string 类型）
	redisKeyPrefix = "session:"

	// cookie的键
	cookieKey = "session-id"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	store, err := redisstore.NewRedisStore(context.TODO(), client)
	if err != nil {
		panic(err)
	}
	defer store.Close()
	// 自定义: Redis中的key的前缀
	store.KeyPrefix(redisKeyPrefix)
	// 自定义: cookie的配置
	store.Options(sessions.Options{
		HttpOnly: false,
		Secure:   false,
		//Path: "/path",
		//Domain: "localhost",
		MaxAge: 1800, // 只有 >=0 的情况下，才会将数据写到Redis中
	})
	// 自定义: cookie的value、Redis中的key的后半部分
	store.KeyGen(func() (string, error) {
		return idKit.NewULID(), nil
	})

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		/* (1) 获取session */
		session, err := store.Get(ctx.Request, cookieKey)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		/* (2) 从session中读取或存储数据 */
		count, err := intKit.ParseToInt(session.Values["count"])
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		count++
		session.Values["count"] = count

		/* (3) 保存session数据，本质上是将内存中的数据持久化到存储介质中（序列化并写到Redis中；会重置key的TTL） */
		if err := session.Save(ctx.Request, ctx.Writer); err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ctx.String(http.StatusOK, "count: [%d]", count)
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
