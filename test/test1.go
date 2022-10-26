package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// 允许跨域
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	id := ulid.Make()
	fmt.Println(id.String())

	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	if err != nil {
		panic(err)
	}
	fmt.Println(id.String())

	//t := time.Now().UTC()
	//entropy := rand.New(rand.NewSource(t.UnixNano()))
	//id := ulid.MustNew(ulid.Timestamp(t), entropy)
	//// id: 01G902ZSM96WV5D5DC5WFHF8WY length: 26
	//fmt.Println("id:", id.String(), "length:", len(id.String()))
}

func main1() {
	r := gin.Default()
	r.GET("/ping", ping)
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

func ping(ctx *gin.Context) {
	// 升级get请求为webSocket协议（如果返回的err != nil，说明websocket连接已经建立成功了）
	ws, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()

	for {
		// 接收前端发送的数据
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			logrus.Error(err)
			break
		}

		message := string(p)
		var response string
		switch message {
		case "ping":
			response = "pong"

			for i := 0; i < 100; i++ {
				text := strconv.Itoa(i)

				go func(text string) {
					if err := ws.WriteMessage(messageType, []byte(text)); err != nil {
						logrus.Error(err)
					}
				}(text)
			}
		default:
			response = fmt.Sprintf("receive message: [%s]", message)
		}

		// 向前端推送数据
		if err := ws.WriteMessage(messageType, []byte(response)); err != nil {
			logrus.Error(err)
			break
		}
	}
}
