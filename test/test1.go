package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// 允许跨域
var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Second * 6,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	_ = r.Run(":8080")
}

func ping(ctx *gin.Context) {
	// 升级get请求为webSocket协议（如果返此处回的err != nil，说明websocket连接已经建立成功了）
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		// 读取前端发来的消息
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			logrus.Errorf("读取数据失败，error: %v", err)
			break
		}

		message := string(msgData)
		var responseText string
		switch message {
		case "ping":
			responseText = "pong"
		default:
			responseText = fmt.Sprintf("receive message: [%d, %s].", msgType, message)
		}

		// 推送消息给前端
		err = conn.WriteMessage(websocket.TextMessage, []byte(responseText))
		if err != nil {
			logrus.Errorf("发送数据失败，error: %v", err)
			break
		}
	}
}
