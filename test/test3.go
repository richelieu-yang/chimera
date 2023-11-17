package main

import (
	"crypto/tls"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaodeKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

var client *gaodeKit.Client

func main() {
	var err error
	client, err = gaodeKit.NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		logrus.Fatal(err)
	}

	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式
	httpClient := bot.Caller.Client.HTTPClient()
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// 注册消息处理函数
	bot.MessageHandler = func(msg *openwechat.Message) {

		msg.IsSendBySelf()

		if msg.IsText() && msg.Content == "ping" {
			_, err := msg.ReplyText("pong")
			if err != nil {
				logrus.WithError(err).Error("Fail to pong")
			} else {
				logrus.Info("Manager to pong")
			}
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(friends)

	results := friends.SearchByRemarkName(1, "狗蛋妈")
	if results.Count() == 0 {
		logrus.Fatal("results.Count() == 0")
	}
	wife := results.First()

	go func() {
		send(wife)
	}()

	if err := bot.Block(); err != nil {
		logrus.Fatal(err)
	}
}

// 发送问候
func send(f *openwechat.Friend) {
	cron := cronKit.NewCron()

	_, err := cron.AddFunc("0 30 6 * * *", greet(f, "宝贝，早安！该起床了。"))
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 11 * * *", greet(f, "宝贝，午安！"))
	if err != nil {
		logrus.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 21 * * *", greet(f, "宝贝，晚安！该困告了。"))
	if err != nil {
		logrus.Fatal(err)
	}

	cron.Start()
}

func greet(f *openwechat.Friend, text string) func() {
	return func() {
		_, err := f.SendText(text)
		if err != nil {
			logrus.WithError(err).WithField("text", text).Error("Fail to greet.")
		} else {
			logrus.Info("Manager to greet.")
		}

	}
}
