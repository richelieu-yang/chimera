package redisKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/confKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
	"time"
)

// 测试: Redis客户端通用订阅与发布.
func TestClient_SubscribeAndPublish(t *testing.T) {

}

// 测试: Redis客户端监听某一db中key的超时
/*
!!!: 需要先配置Redis.
*/
func TestClient_SubscribeAndPublish1(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		logrus.Infof("wd: %s", wd)
	}
	logrus.Infof("wd: %s", wd)

	type config struct {
		Redis Config `json:"redis"`
	}
	c := &config{}
	path := "chimera-lib/config.yaml"
	confKit.MustLoad(path, c)
	MustSetUp(c.Redis)
	client, err := GetClient()
	if err != nil {
		panic(err)
	}
	client = client

	/* 方法1 */
	//go func() {
	//	pubSub := client.Subscribe(context.TODO(), "__keyevent@0__:expired")
	//	defer pubSub.Close()
	//
	//	for {
	//		msg, err := pubSub.ReceiveMessage(context.TODO())
	//		if err != nil {
	//			panic(err)
	//		}
	//		logrus.WithFields(logrus.Fields{
	//			"channel": msg.Channel,
	//			"payLoad": msg.Payload, // 过期的键（key）
	//		}).Info("Receive a message.")
	//	}
	//}()
	/* 方法2 */
	go func() {
		pubSub := client.Subscribe(context.TODO(), "__keyevent@0__:expired")
		defer pubSub.Close()

		ch := pubSub.Channel()
		for msg := range ch {
			logrus.WithFields(logrus.Fields{
				"channel": msg.Channel,
				"payLoad": msg.Payload, // 过期的键（key）
			}).Info("Receive a message.")
		}
	}()

	go func() {
		flag, err := client.Set(context.TODO(), "a", "aaa", time.Second*3)
		if err != nil {
			panic(err)
		}
		logrus.Infof("flag: %t", flag)
	}()

	time.Sleep(time.Second * 5)
}
