package rocketmq5Kit

import (
	"context"
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/richelieu42/chimera/src/idKit"
	"github.com/richelieu42/chimera/src/jsonKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	// producerTimeout 单次推送消息的timeout
	producerTimeout = time.Millisecond * 500
	// verifyTimeout 验证的最长timeout
	verifyTimeout = time.Second * 6
)

// VerifyEndpoint 测试RocketMQ5服务是否启动正常.
/*
@param endpoint 用于测试的RocketMQ5服务的endpoint
@param topic 	用于测试的topic（理论上，此topic仅用于测试，不能同时用于业务，因为测试发的消息无意义）
@return 如果为nil，说明 RocketMQ5服务 正常启动
*/
func VerifyEndpoint(endpoint, topic string) error {
	if strKit.IsEmpty(endpoint) {
		return errorKit.Simple("param endpoint is empty")
	}
	if strKit.IsEmpty(topic) {
		return errorKit.Simple("param topic is empty")
	}

	/* logger */
	tempDir, err := pathKit.GetChimeraTempDir()
	if err != nil {
		return err
	}
	logName := fmt.Sprintf("rocketmq5_%s_%s.log", topic, idKit.NewULID())
	logPath := pathKit.Join(tempDir, logName)
	logger, err := logrusKit.NewFileLogger(logPath, nil, logrus.DebugLevel, false)
	if err != nil {
		return err
	}
	defer func() {
		_ = logrusKit.DisposeLogger(logger)
	}()
	logger.Infof("endpoint: [%s].", endpoint)
	logger.Infof("topic: [%s].", topic)

	/* texts */
	timeStr := timeKit.FormatCurrentTime()
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$3"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$4"),
		fmt.Sprintf("%s_%s_%s", ulid, timeStr, "$5"),
	}
	json, err := jsonKit.MarshalToStringWithIndent(texts)
	if err != nil {
		return err
	}
	logger.Infof("texts: %v.", json)

	mqLogConfig := &LogConfig{
		ToConsole: false,
		LogDir:    tempDir,
		LogName:   logName,
	}
	mqConfig := &rmq_client.Config{
		Endpoint: endpoint,
	}

	/* consumer */
	consumer, err := NewSimpleConsumer(mqLogConfig, mqConfig, fmt.Sprintf("%s-%s", topic, idKit.NewULID()), topic, "*")
	if err != nil {
		return errorKit.Wrap(err, "fail to new consumer")
	}
	if err := consumer.Start(); err != nil {
		return errorKit.Wrap(err, "fail to start consumer with topic(%s)", topic)
	}
	defer consumer.GracefulStop()

	/* producer */
	producer, err := NewProducer(mqLogConfig, mqConfig)
	if err != nil {
		return errorKit.Wrap(err, "fail to new producer")
	}
	if err := producer.Start(); err != nil {
		return errorKit.Wrap(err, "fail to start producer")
	}
	defer producer.GracefulStop()

	// test
	//logrus.Infof("logPath: [%s].", logPath)

	producerCh := make(chan error, 1)
	consumerCh := make(chan error, 1)
	consumerCtx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	/* consumer works */
	go func() {
		defer func() {
			logger.Info("[Consumer] Goroutine ends.")
		}()

		duplicate := sliceKit.Copy(texts)
		for {
			time.Sleep(time.Millisecond * 100)

			mvs, err := consumer.Receive(consumerCtx, MaxMessageNum, InvisibleDuration)
			if err != nil {
				// 时间到
				if s, ok := status.FromError(err); ok {
					switch s.Code() {
					case codes.Canceled:
						consumerCh <- err
						return
						//case codes.Canceled:
						//	/* 提前结束（被取消） */
						//	break LOOP
						//case codes.DeadlineExceeded:
						//	/* 超时结束 */
						//	consumerErr = errorKit.Simple("consumer fails to receive all messages(count: %d) within timeout(%s), missing(%d)", len(texts), verifyTimeout.String(), len(text1))
						//	break LOOP
					}
				}

				if errRpcStatus, ok := rmq_client.AsErrRpcStatus(err); ok {
					switch errRpcStatus.Code {
					case int32(v2.Code_MESSAGE_NOT_FOUND):
						// 没有新消息
					default:
						logger.WithFields(logrus.Fields{
							"code":  errRpcStatus.Code,
							"error": err.Error(),
						}).Warn("[Consumer] Fail to receive.")
					}
				} else {
					logger.WithFields(logrus.Fields{
						"error": err.Error(),
					}).Warn("[Consumer] Fail to receive.")
				}
				continue
			}

			// ack message
			for _, mv := range mvs {
				text := string(mv.GetBody())

				err := consumer.Ack(context.TODO(), mv)
				if err != nil {
					logger.WithFields(logrus.Fields{
						"text":  text,
						"error": err.Error(),
					}).Error("[Consumer] Fail to ack.")
					continue
				}

				var ok bool
				duplicate, ok = sliceKit.Remove(duplicate, text)
				left := len(duplicate)
				logger.WithFields(logrus.Fields{
					"valid": ok,
					"left":  left,
					"text":  text,
				}).Info("[CONSUMER] Receive and ack a message.")
				if left == 0 {
					// 成功收到所有预期消息
					logger.Info("[CONSUMER] Receive and ack all messages.")

					consumerCh <- nil
					return
				}
			}
		}
	}()

	/* producer works */
	go func() {
		defer func() {
			logger.Info("[Producer] Goroutine ends.")
		}()

		// 等一会，以确保: producer发消息时，consumer已经开始收消息了（以免丢失消息: 明明producer发了，但consumer却没收到）
		time.Sleep(time.Second)

		for _, text := range texts {
			msg := &rmq_client.Message{
				Topic: topic,
				Body:  []byte(text),
			}
			ctx, _ := context.WithTimeout(context.TODO(), producerTimeout)
			_, err := producer.Send(ctx, msg)
			if err != nil {
				err = errorKit.Wrap(err, "[Producer] Fail to send message(%s).", text)
				producerCh <- err
				return
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[Producer] Send a message successfully.")
		}
		logger.Info("[Producer] Send all messages successfully.")
	}()

	select {
	case err = <-producerCh:
	case err = <-consumerCh:
		// 此处err可能为nil（说明验证通过）
	case <-time.After(verifyTimeout):
		err = errorKit.Simple("fail to pass validation within timeout(%v)", verifyTimeout)
	}

	if err != nil {
		err = errorKit.Wrap(err, "log path: [%s]", logPath)
		logger.Errorf("%+v", err)
		return err
	}
	return nil
}
