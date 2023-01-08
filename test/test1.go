package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore/v8"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	var store *redisstore.RedisStore
	// 存储到Redis中的key的前缀（value为类型）
	keyPrefix := "session:"
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   10,
	})
	store, err := redisstore.NewRedisStore(context.TODO(), client)
	if err != nil {
		panic(err)
	}
	store.KeyPrefix(keyPrefix)
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
		cookieKey := "session-id"
		session, err := store.Get(ctx.Request, cookieKey)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		/* (2) 从session中读取或存储数据 */
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

	// Request y writer for testing
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	w := httptest.NewRecorder()

	// Get session
	session, err := store.Get(req, "session-key")
	if err != nil {
		log.Fatal("failed getting session: ", err)
	}

	// Add a value
	session.Values["foo"] = "bar"

	// Save session
	if err = sessions.Save(req, w); err != nil {
		log.Fatal("failed saving session: ", err)
	}

	// Delete session (MaxAge <= 0)
	session.Options.MaxAge = -1
	if err = sessions.Save(req, w); err != nil {
		log.Fatal("failed deleting session: ", err)
	}
}
