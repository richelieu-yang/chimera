package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/mapKit"
	"net/http"
	"sync/atomic"
	"time"
)

var (
	// 存储到Redis中的key的前缀（value为类型）
	redisKeyPrefix = "session:"

	// cookie的键
	cookieKey = "session-id"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   10,
	})
	store, err := redisstore.NewRedisStore(context.TODO(), client)
	if err != nil {
		panic(err)
	}
	store.KeyPrefix(redisKeyPrefix)
	store.Options(sessions.Options{
		HttpOnly: false,
		Secure:   false,
		//Path: "/path",
		//Domain: "localhost",
		MaxAge: 3600,
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
		if !mapKit.Contains(session.Values, "count") {
			session.Values["count"] = 0
		}
		count, err := intKit.ParseToInt32(session.Values["count"])
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		atomic.AddInt32(&count, 1)

		timeStr := time.Now().String()
		session.Values["time"] = timeStr

		/* (3) 保存session数据。本质上是将内存中的数据持久化到存储介质中 */
		if err := session.Save(ctx.Request, ctx.Writer); err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		ctx.String(http.StatusOK, timeStr)
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
