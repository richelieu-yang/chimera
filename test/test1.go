package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// 允许跨域
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
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
