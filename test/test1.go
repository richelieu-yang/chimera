package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
)

var (

	sessions.

	store *sessions.RedisStore
)

func init() {
	// 创建一个Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器地址
		Password: "",               // Redis服务器密码（如果有）
		DB:       0,                // 使用默认的数据库
	})

	// 创建一个Redis存储引擎
	store = sessions.NewRedisStore(client, time.Hour)

	// 设置存储引擎的密钥
	store.KeyPrefix("session:")
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400, // 会话过期时间（秒）
		HttpOnly: true,
	})
}

func main() {
	// 配置路由处理程序
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)

	// 启动HTTP服务器
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	// 获取会话对象
	session, _ := store.Get(r, "session-name")

	// 设置会话值
	session.Values["username"] = "John Doe"
	session.Values["email"] = "johndoe@example.com"

	// 保存会话
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Session has been set")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// 获取会话对象
	session, _ := store.Get(r, "session-name")

	// 从会话中检索值
	username := session.Values["username"]
	email := session.Values["email"]

	// 将会话值发送给客户端
	fmt.Fprintf(w, "Username: %s\nEmail: %s\n", username, email)
}
