package main

import (
	"crypto/tls"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/thirdParty/gaodeKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var client *gaodeKit.Client
var logger *logrus.Logger

func init() {
	logrusKit.MustSetUp(nil)
	logrusKit.DisableQuote(nil)

	var err error

	path := fmt.Sprintf("%s.log", timeKit.FormatCurrent(timeKit.FormatFileName))
	file, err := fileKit.Create(path)
	if err != nil {
		logrus.Fatal(err)
	}
	out := ioKit.MultiWriter(file, os.Stdout)
	logger = logrusKit.NewLogger(logrusKit.WithDisableQuote(true), logrusKit.WithOutput(out))

	client, err = gaodeKit.NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		logrus.Fatal(err)
	}
}

func main() {
	var err error
	client, err = gaodeKit.NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		logger.Fatal(err)
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
		if msg.IsText() && msg.Content == "ping" {
			_, err := msg.ReplyText("pong")
			if err != nil {
				logger.WithError(err).Error("Fail to pong")
			} else {
				logger.Info("Manager to pong")
			}
		}
	}
	// 注册登陆二维码回调
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	// 登陆
	if err := bot.Login(); err != nil {
		logger.Fatal(err)
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		logger.Fatal(err)
	}

	// 获取所有的好友
	friends, err := self.Friends()
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(friends)

	nickName := "软心姑娘sss"
	results := friends.SearchByNickName(1, nickName)
	//results := friends.SearchByRemarkName(1, "狗蛋妈")
	if results.Count() == 0 {
		logger.WithField("nickName", nickName).Fatal("Fail to get friend.")
	}
	wife := results.First()
	logger.WithField("nickName", nickName).Info("Manager to get friend.")

	go func() {
		start(wife)
	}()

	if err := bot.Block(); err != nil {
		logger.Fatal(err)
	}
}

// 发送问候
func start(f *openwechat.Friend) {
	cron := cronKit.NewCron()

	_, err := cron.AddFunc("0 30 6 * * *", task(f, "宝贝，早安！该起床了。"))
	if err != nil {
		logger.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 11 * * *", task(f, "宝贝，午安！"))
	if err != nil {
		logger.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 17 * * *", task(f, "宝贝，我下班喽。"))
	if err != nil {
		logger.Fatal(err)
	}

	_, err = cron.AddFunc("0 30 21 * * *", task(f, "宝贝，晚安！该困告了。"))
	if err != nil {
		logger.Fatal(err)
	}

	cron.Start()
}

func task(f *openwechat.Friend, greetText string) func() {
	return func() {
		// greet
		logger.Infof("greetText: %s", greetText)
		_, err := f.SendText(greetText)
		if err != nil {
			logger.WithError(err).Error("Fail to greet.")
		} else {
			logger.Info("Manager to greet.")
		}

		// weather
		weather := getWeather()
		logger.Infof("weather: %s", weather)
		_, err = f.SendText(weather)
		if err != nil {
			logger.WithError(err).Error("Fail to greet.")
		} else {
			logger.Info("Manager to send weather.")
		}
	}
}

func getWeather() string {
	city := "320200"
	live, err := client.GetLive(city)
	if err != nil {
		panic(err)
	}
	cast, err := client.GetTodayCast(city)
	if err != nil {
		panic(err)
	}

	text := fmt.Sprintf("您好，现在是[%s]，[%s %s]目前气温：%s，白天气温：%s，夜晚气温：%s.",
		live.ReportTime,
		live.Province,
		live.City,
		fmt.Sprintf("%s°C(%s)", live.Temperature, live.Weather),
		fmt.Sprintf("%s°C(%s)", cast.DayTemp, cast.DayWeather),
		fmt.Sprintf("%s°C(%s)", cast.NightTemp, cast.NightWeather))
	return text
}
