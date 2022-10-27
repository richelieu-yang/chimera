package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/http/httpKit"
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
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()

	if err := engine.SetTrustedProxies([]string{"172.19.10.12"}); err != nil {
		panic(err)
	}

	//engine.TrustedPlatform

	engine.GET("/ping", ping)
	_ = engine.Run(":8080")
}

func ping(ctx *gin.Context) {

	ctx.ClientIP()

	// 先判断是不是websocket请求
	if !websocket.IsWebSocketUpgrade(ctx.Request) {
		// 普通http请求
		ctx.String(http.StatusOK, `request(method: %s, url: %s, Connection: %s, Upgrade: %s) isn't a websocket request`,
			ctx.Request.Method, httpKit.GetRequestUrl(ctx.Request), ctx.Request.Header["Connection"], ctx.Request.Header["Upgrade"])
		return
	}

	// 升级请求为WebSocket协议（如果返此处回的err != nil，说明websocket连接已经建立成功了，即握手成功）
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
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
