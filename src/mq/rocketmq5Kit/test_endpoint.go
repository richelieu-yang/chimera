package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/idKit"
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	producerTimeout = time.Millisecond * 300
	consumerTimeout = time.Second * 3
)

// TestEndpoint 测试RocketMQ5服务是否启动正常.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func TestEndpoint(endpoint, topic string) (finalErr error) {
	/* texts */
	timeStr := timeKit.FormatCurrentTime()
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$3"),
	}

	config := &rmq_client.Config{
		Endpoint: endpoint,
	}

	/* log */
	tempDir, err := pathKit.GetTempDirOfGoScales()
	if err != nil {
		return errorKit.Wrap(err, "fail to get temporary directory")
	}
	logName := fmt.Sprintf("rocketmq5_test_%s_%s_%s.log", endpoint, topic, timeStr)
	// 文件名不支持":"（无论是Mac还是Windows）
	logName = strKit.ReplaceAll(logName, ":", "：")
	logPath := pathKit.Join(tempDir, logName)
	var logConfig = &LogConfig{
		ToConsole: false,
		LogDir:    tempDir,
		LogName:   logName,
	}
	logger, err := logrusKit.NewFileLogger(logPath, nil, logrus.DebugLevel, false)
	if err != nil {
		return errorKit.Wrap(err, "fail to new logger")
	}
	logger.Infof("endpoint: [%s].", endpoint)
	logger.Infof("topic: [%s].", topic)
	defer func() {
		_ = logrusKit.DisposeLogger(logger)

		if finalErr != nil {
			finalErr = errorKit.Wrap(finalErr, "log path: %s", logPath)
		} else {
			// 通过测试的话，删除日志文件
			_ = fileKit.Delete(logPath)
		}
	}()

	/* Producer */
	producer, err := NewProducer(logConfig, config)
	if err != nil {
		finalErr = errorKit.Wrap(err, "fail to new producer")
		return
	}
	if err := producer.Start(); err != nil {
		finalErr = errorKit.Wrap(err, "fail to start producer")
		return
	}
	defer producer.GracefulStop()

	/* Consumer */
	consumerGroup := fmt.Sprintf("simpleConsumer-group-%s-%s", topic, idKit.NewULID())
	simpleConsumer, err := NewSimpleConsumer(logConfig, config, consumerGroup, topic, "*")
	if err != nil {
		finalErr = errorKit.Wrap(err, "fail to new simple consumer")
		return
	}
	if err := simpleConsumer.Start(); err != nil {
		finalErr = errorKit.Wrap(err, "fail to start simple consumer")
		return
	}
	defer simpleConsumer.GracefulStop()

	consumerCtx, cancel := context.WithTimeout(context.TODO(), consumerTimeout)
	defer cancel()

	/* 启动Producer */
	var producerErr error
	go func() {
		defer func() {
			logger.Info("[Producer] Goroutine ends.")
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

			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[Producer] Send message successfully.")
		}

		if err != nil {
			producerErr = err
			cancel()
			logger.Infof("[Producer] Invoke cancel() with error: %+v", err)
			return
		}
		logger.WithFields(logrus.Fields{
			"count": len(texts),
		}).Info("[Producer] Send all messages successfully.")

	}()

	/* 启动Consumer */
	var consumerErr error
	// 拷贝，不修改texts
	var text1 = sliceKit.Copy(texts)
LOOP:
	for {
		mvs, err := simpleConsumer.Receive(consumerCtx, MaxMessageNum, InvisibleDuration)
		if err != nil {
			// 时间到
			if s, ok := status.FromError(err); ok {
				switch s.Code() {
				case codes.Canceled:
					/* 提前结束（被取消） */
					break LOOP
				case codes.DeadlineExceeded:
					/* 超时结束 */
					consumerErr = errorKit.Simple("consumer fails to receive all messages(count: %d) within timeout(%s), missing(%d)", len(texts), consumerTimeout.String(), len(text1))
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
			consumerErr = errorKit.Wrap(err, "consumer meets an unprocessed error")
			break LOOP
		}

		// ack message
		for _, mv := range mvs {
			err := simpleConsumer.Ack(context.TODO(), mv)
			if err != nil {
				/* 提前结束（确认消息） */
				consumerErr = errorKit.Wrap(err, "consumer fails to ack message")
				break LOOP
			}
			text := string(mv.GetBody())

			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[CONSUMER] Receive a message.")

			var ok bool
			text1, ok = sliceKit.Remove(text1, text)
			if ok {
				if sliceKit.IsEmpty(text1) {
					/* 提前结束（已收到所有消息） */
					consumerErr = nil
					break LOOP
				}
			}
		}

		// 每次收消息间有一定的间隔
		time.Sleep(time.Millisecond * 100)
	}

	if producerErr != nil {
		logger.Errorf("Fail to pass test, error: %+v", producerErr)
		finalErr = producerErr
		return
	}
	if consumerErr != nil {
		logger.Errorf("Fail to pass test, error: %+v", consumerErr)
		finalErr = consumerErr
		return
	}
	logger.Info("Pass test.")
	return
}
