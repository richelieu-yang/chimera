package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/core/timeKit"
	"github.com/richelieu42/go-scales/src/idKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	producerTimeout = time.Millisecond * 500
	consumerTimeout = time.Second * 10
)

// TestEndpoint 测试RocketMQ5服务是否启动正常.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func TestEndpoint(endpoint, topic string) (finalErr error) {
	if strKit.IsEmpty(endpoint) {
		return errorKit.Simple("param endpoint is empty")
	}
	if strKit.IsEmpty(topic) {
		return errorKit.Simple("param topic is empty")
	}

	/* logger */
	tempDir, err := pathKit.GetTempDirOfGoScales()
	if err != nil {
		return err
	}
	logName := fmt.Sprintf("rocketmq5_%s.log", idKit.NewULID())
	logPath := pathKit.Join(tempDir, logName)
	logger, err := logrusKit.NewFileLogger(logPath, nil, logrus.DebugLevel, false)
	if err != nil {
		return err
	}
	defer logrusKit.DisposeLogger(logger)

	/* texts */
	timeStr := timeKit.FormatCurrentTime()
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$3"),
	}

	/* print */
	logger.Infof("endpoint: [%s].", endpoint)
	logger.Infof("topic: [%s].", topic)
	json, err := jsonKit.MarshalToStringWithIndent(texts)
	if err != nil {
		return err
	}
	logger.Infof("texts: %v.", json)

	//defer func() {
	//	if finalErr != nil {
	//		finalErr = errorKit.Wrap(finalErr, "log path: %s", logPath)
	//	} else {
	//		// 通过测试的话，删除日志文件
	//		_ = fileKit.Delete(logPath)
	//	}
	//}()

	mqLogConfig := &LogConfig{
		ToConsole: false,
		LogDir:    tempDir,
		LogName:   logName,
	}
	mqConfig := &rmq_client.Config{
		Endpoint: endpoint,
	}

	/* producer */
	producer, err := NewProducer(mqLogConfig, mqConfig)
	if err != nil {
		return err
	}
	if err := producer.Start(); err != nil {
		return err
	}
	defer producer.GracefulStop()

	/* consumer */
	consumer, err := NewSimpleConsumer(mqLogConfig, mqConfig, fmt.Sprintf("%s-%s", topic, idKit.NewULID()), topic, "*")
	if err != nil {
		return err
	}
	if err := consumer.Start(); err != nil {
		return err
	}
	defer consumer.GracefulStop()

	consumerCtx, cancel := context.WithTimeout(context.TODO(), consumerTimeout)
	defer cancel()

	/* start producer */
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

	/* start consumer */
	var consumerErr error
	// 拷贝，不修改texts
	var text1 = sliceKit.Copy(texts)
LOOP:
	for {
		mvs, err := consumer.Receive(consumerCtx, MaxMessageNum, InvisibleDuration)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"error": err.Error(),
			}).Warn("fail to receive")

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

			///* 提前结束（未知错误） */
			//consumerErr = errorKit.Wrap(err, "consumer meets an unprocessed error")
			//break LOOP

			continue
		}

		// ack message
		for _, mv := range mvs {
			err := consumer.Ack(context.TODO(), mv)
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
		logger.Errorf("Fail to pass test, producerErr: %+v", producerErr)
		finalErr = producerErr
		return
	}
	if consumerErr != nil {
		logger.Errorf("Fail to pass test, consumerErr: %+v", consumerErr)
		finalErr = consumerErr
		return
	}
	logger.Info("Pass test.")
	return
}
