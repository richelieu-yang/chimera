/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"os"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/apache/rocketmq-clients/golang/v5/credentials"
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5

	// maximum number of messages received at one time
	maxMessageNum int32 = 16

	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
)

func main() {
	logrusKit.MustSetUp(nil)

	//Topic := "test"
	Endpoint := "127.0.0.1:8081"
	AccessKey := ""
	SecretKey := ""
	cg := "ConsumerGroup"

	// log to console
	if err := os.Setenv("mq.consoleAppender.enabled", "true"); err != nil {
		logrus.Fatal(err)
	}
	rmq_client.ResetLogger()

	// In most case, you don't need to create many consumers, singletion pattern is more recommended.
	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: cg,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		rmq_client.WithAwaitDuration(awaitDuration),
		rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
			Topic: rmq_client.SUB_ALL,
		}),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	// start simpleConsumer
	err = simpleConsumer.Start()
	if err != nil {
		logrus.Fatal(err)
	}
	// graceful stop simpleConsumer
	defer simpleConsumer.GracefulStop()

	go func() {
		for {
			mvs, err := simpleConsumer.Receive(context.TODO(), maxMessageNum, invisibleDuration)
			if err != nil {
				logrus.WithError(err).Error("[CONSUMER] fail to receive")
			}
			// ack message
			for _, mv := range mvs {
				if err := simpleConsumer.Ack(context.TODO(), mv); err != nil {
					logrus.WithFields(logrus.Fields{
						"topic": mv.GetTopic(),
						"tag":   mv.GetTag(),
						"msgId": mv.GetMessageId(),
						"text":  string(mv.GetBody()),
						"error": err,
					}).Error("[CONSUMER] fail to ack the message")
				}
				logrus.WithFields(logrus.Fields{
					"msgId": mv.GetMessageId(),
					"text":  string(mv.GetBody()),
				}).Info("[CONSUMER] receive a message")
			}

			logrus.Debug("[CONSUMER] wait starts")
			time.Sleep(time.Second)
			logrus.Debug("[CONSUMER] wait ends")
		}
	}()

	select {}
}
