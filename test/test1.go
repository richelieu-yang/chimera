package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/idKit"
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
	id := idKit.NewULID()
	fmt.Println(id)
	fmt.Println(len(id))

	//engine := gin.Default()
	//
	////engine.GET("/ping", ping)
	//engine.Any("/ping", wsKit.GinHandler)
	//
	//_ = engine.Run(":8080")
}

func ping(ctx *gin.Context) {
	// 先判断是不是websocket请求
	if !websocket.IsWebSocketUpgrade(ctx.Request) {
		ctx.String(http.StatusOK, "request(method: %s, Connection: %s, Upgrade: %s) isn't a websocket request",
			ctx.Request.Method, ctx.Request.Header["Connection"], ctx.Request.Header["Upgrade"])
		return
	}

	// 升级get请求为webSocket协议（如果返此处回的err != nil，说明websocket连接已经建立成功了）
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
	if err != nil {
		logrus.Errorf("fail to upgrade, error: %v", err)
		return
	}
	defer conn.Close()

	// 监听连接的断开
	conn.SetCloseHandler(func(code int, text string) error {
		logrus.Warnf("connection is closed with code(%d) and text(%s)", code, text)

		// 默认的close handler
		message := websocket.FormatCloseMessage(code, text)
		_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

LOOP:
	for {
		// 读取前端发来的消息
		msgType, msgData, err := conn.ReadMessage()
		if err != nil {
			logrus.Errorf("fail to read text，error: %v", err)
			break
		}

		text := string(msgData)
		var responseText string
		switch text {
		case "ping":
			responseText = "pong"
		case "close":
			if err := conn.Close(); err != nil {
				logrus.Errorf("fail to close connection，error: %v", err)
			}
			break LOOP
		default:
			responseText = fmt.Sprintf("receive message(type: %d, text: %s)", msgType, text)
		}

		// 推送消息给前端
		err = conn.WriteMessage(websocket.TextMessage, []byte(responseText))
		if err != nil {
			logrus.Errorf("fail to write text，error: %v", err)
			break
		}
	}
}
