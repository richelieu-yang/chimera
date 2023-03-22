package pulsarKit

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/richelieu42/chimera/src/core/pathKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/richelieu42/chimera/src/idKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/operationKit"
	"github.com/sirupsen/logrus"
	"time"
)

// verify 简单地验证 Pulsar服务 是否启动成功
func verify(verifyConfig *VerifyConfig) (err error) {
	dir, _ := pathKit.GetChimeraTempDir()
	timeStr := timeKit.FormatCurrentTime(timeKit.FormatFileName)
	consumerLogPath := pathKit.Join(dir, fmt.Sprintf("pulsar_verify_consumer_%s.log", timeStr))
	producerLogPath := pathKit.Join(dir, fmt.Sprintf("pulsar_verify_producer_%s.log", timeStr))
	defer func() {
		if err == nil {
			// 验证成功，删掉客户端日志文件
			_ = fileKit.Delete(consumerLogPath)
			_ = fileKit.Delete(producerLogPath)
		}
	}()

	err = _verify(verifyConfig, consumerLogPath, producerLogPath)
	return
}

func _verify(verifyConfig *VerifyConfig, consumerLogPath, producerLogPath string) error {
	// 接受消息的时限
	var receiveTimeLimit = time.Second * 10
	// 单次发送消息的超时时间
	var sendTimeout = time.Second

	if verifyConfig == nil {
		// 不验证
		return nil
	}
	topic := verifyConfig.Topic
	if strKit.IsEmpty(topic) || strKit.IsBlank(topic) {
		// 不验证
		return nil
	}
	// 是否打印日志到控制台？
	printFlag := verifyConfig.Print
	logger := logrusKit.NewLogger(nil, operationKit.Ternary(printFlag, logrus.DebugLevel, logrus.PanicLevel))
	logger.Infof("[Consumer] log path: [%s].", consumerLogPath)
	logger.Infof("[Producer] log path: [%s].", producerLogPath)

	timeStr := timeKit.FormatCurrentTime()
	ulid := idKit.NewULID()
	texts := []string{
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$3"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$4"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$5"),
	}

	consumer, err := NewConsumer(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: ulid,
		Type:             pulsar.Exclusive,
	}, consumerLogPath)
	if err != nil {
		return err
	}
	defer consumer.Close()
	producer, err := NewProducer(pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	}, producerLogPath)
	if err != nil {
		return err
	}
	defer producer.Close()

	var ch = make(chan struct{}, 1)
	var consumerErrCh = make(chan error, 1)
	var producerErrCh = make(chan error, 1)

	/* consumer */
	consumerCtx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go func() {
		defer func() {
			logger.Info("[Consumer] goroutine ends")
		}()

		s := sliceKit.Copy(texts)

		for {
			msg, err := consumer.Receive(consumerCtx)
			if err != nil {
				err = errorKit.Wrap(err, "fail to receive")
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Consumer] fail to receive")
				consumerErrCh <- err
				break
			}
			if err := consumer.Ack(msg); err != nil {
				err = errorKit.Wrap(err, "fail to ack")
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Consumer] fail to ack")
				consumerErrCh <- err
				break
			}

			var ok bool
			text := string(msg.Payload())
			s, ok = sliceKit.Remove(s, text)
			left := len(s)
			logger.WithFields(logrus.Fields{
				"left":  left,
				"valid": ok,
				"text":  text,
			}).Info("[Consumer] receive a message")

			if ok && left == 0 {
				logger.Info("[Consumer] receive all messages!")
				ch <- struct{}{}
				break
			}
		}
	}()

	/* producer */
	go func() {
		defer func() {
			logger.Info("[Producer] goroutine ends")
		}()

		for _, text := range texts {
			pMsg := &pulsar.ProducerMessage{
				Payload: []byte(text),
			}
			err := func() error {
				ctx, cancel := context.WithTimeout(context.TODO(), sendTimeout)
				defer cancel()
				_, err := producer.Send(ctx, pMsg)
				return err
			}()
			if err != nil {
				err = errorKit.Wrap(err, "fail to send")
				logger.WithFields(logrus.Fields{
					"text":  text,
					"error": err.Error(),
				}).Error("[Producer] fail to send")
				producerErrCh <- err
				break
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[Producer] succeeded to send")
		}
	}()

	select {
	case <-ch:
		return nil
	case err := <-producerErrCh:
		return err
	case err := <-consumerErrCh:
		return err
	case <-time.After(receiveTimeLimit):
		return errorKit.Simple("fail to get all messages within time limit(%s)", receiveTimeLimit)
	}
}
