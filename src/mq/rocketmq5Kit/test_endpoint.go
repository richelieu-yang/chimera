package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/idKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

// TestEndpoint 测试RocketMQ5服务是否启动正常.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func TestEndpoint(endpoint, topic string) error {
	producerTimeout := time.Millisecond * 300
	consumerTimeout := time.Second * 10

	timeStr := timeKit.FormatCurrentTime()
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$3"),
	}
	textsCount := len(texts)

	config := &rmq_client.Config{
		Endpoint: endpoint,
	}
	var logConfig *LogConfig = nil

	/* Producer */
	producer, err := NewProducer(logConfig, config)
	if err != nil {
		return errorKit.Wrap(err, "fail to new producer")
	}
	if err := producer.Start(); err != nil {
		return errorKit.Wrap(err, "fail to start producer")
	}
	defer producer.GracefulStop()

	/* Consumer */
	consumerGroup := fmt.Sprintf("simpleConsumer-group-%s-%s", topic, idKit.NewULID())
	simpleConsumer, err := NewSimpleConsumer(logConfig, config, consumerGroup, topic, "*")
	if err != nil {
		return errorKit.Wrap(err, "fail to new simple consumer")
	}
	if err := simpleConsumer.Start(); err != nil {
		return errorKit.Wrap(err, "fail to start simple consumer")
	}
	defer simpleConsumer.GracefulStop()

	consumerCtx, cancel := context.WithTimeout(context.TODO(), consumerTimeout)
	defer cancel()
	var e error

	/* 启动Producer */
	go func() {
		defer func() {
			logrus.Info("[Producer] Goroutine ends.")
		}()

		// 等100ms，以确保: producer发消息时，consumer已经开始收消息了（以免丢失消息: 明明producer发了，但consumer却没收到）
		time.Sleep(time.Millisecond * 100)

		var err error
		for _, text := range texts {
			msg := &rmq_client.Message{
				Topic: topic,
				Body:  []byte(text),
			}
			ctx, _ := context.WithTimeout(context.TODO(), producerTimeout)
			_, err = producer.Send(ctx, msg)
			if err != nil {
				err = errorKit.Wrap(err, "fail to send message")
				break
			}

			// test
			logrus.Debugf("[SEND] %s", text)
		}
		if err != nil {
			e = err
			cancel()
			return
		}
	}()

	/* 启动Consumer */
LOOP:
	for {
		mvs, err := simpleConsumer.Receive(consumerCtx, MaxMessageNum, InvisibleDuration)
		if err != nil {
			// 时间到
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.Canceled:
					/* 提前结束（被取消） */
					if e == nil {
						e = errorKit.Simple("consumer is canceled")
					}
					break LOOP
				case codes.DeadlineExceeded:
					/* 超时结束 */
					e = errorKit.Simple("consumer fails to receive all messages(count: %d) with timeout(%s) and missing(%d)", textsCount, consumerTimeout.String(), len(texts))
					break LOOP
				}
			}

			if errRpcStatus, ok := rmq_client.AsErrRpcStatus(err); ok {
				switch errRpcStatus.Code {
				case int32(v2.Code_MESSAGE_NOT_FOUND):
					// 每次收消息间有一定的间隔
					time.Sleep(time.Millisecond * 100)
					continue
				}
			}

			/* 提前结束（未知错误） */
			e = errorKit.Wrap(err, "consumer meets an unprocessed error")
			break LOOP
		}

		// ack message
		for _, mv := range mvs {
			err := simpleConsumer.Ack(context.TODO(), mv)
			if err != nil {
				/* 提前结束（确认消息） */
				e = errorKit.Wrap(err, "consumer fails to ack message")
				break LOOP
			}
			text := string(mv.GetBody())

			// test
			logrus.Debugf("[RECEIVE] %s", text)

			var ok bool
			texts, ok = sliceKit.Remove(texts, text)
			if ok {
				if sliceKit.IsEmpty(texts) {
					/* 提前结束（已收到所有消息） */
					e = nil
					break LOOP
				}
			}
		}

		// 每次收消息间有一定的间隔
		time.Sleep(time.Millisecond * 100)
	}

	return e
}
