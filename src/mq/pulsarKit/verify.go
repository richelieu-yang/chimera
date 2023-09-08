package pulsarKit

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/operationKit"
	"github.com/richelieu-yang/chimera/v2/src/randomKit"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	// connectTimeout 创建Consumer（或Producer）的超时时间
	connectTimeout = time.Second * 10

	// receiveTimeout 接受消息的超时时间
	receiveTimeout = time.Second * 10

	// sendTimeout 单次发送消息的超时时间
	sendTimeout = time.Second
)

// verify 简单地验证 Pulsar服务 是否启动成功
/*
PS:
可能失败的原因：
（1）pulsar的进程在，但启动报错（存储空间爆了）
*/
func verify(verifyConfig VerifyConfig, tmpDirPath string) (err error) {
	if strKit.IsBlank(verifyConfig.Topic) {
		// 不验证
		return nil
	}

	ulid := idKit.NewULID()
	topic := verifyConfig.Topic

	if strKit.IsEmpty(tmpDirPath) {
		tmpDirPath, err = pathKit.GetUniqueTempDir()
		if err != nil {
			return err
		}
	} else {
		if err := fileKit.AssertExistAndIsDir(tmpDirPath); err != nil {
			return err
		}
	}
	timeStr := timeKit.FormatCurrent(timeKit.FormatFileName)
	consumerLogPath := pathKit.Join(tmpDirPath, fmt.Sprintf("pulsar_verify_consumer_%s_%s.log", timeStr, ulid))
	producerLogPath := pathKit.Join(tmpDirPath, fmt.Sprintf("pulsar_verify_producer_%s_%s.log", timeStr, ulid))

	// 是否打印日志到控制台？
	printFlag := verifyConfig.Print
	level := operationKit.Ternary(printFlag, logrus.DebugLevel, logrus.PanicLevel)
	cLogger := logrusKit.NewLogger(logrusKit.WithLevel(level))
	cLogger.Infof("[Verify] consumerLogPath: [%s].", consumerLogPath)
	cLogger.Infof("[Verify] producerLogPath: [%s].", producerLogPath)

	defer func() {
		if err == nil {
			// 验证成功的情况下，删掉客户端日志文件
			if err := fileKit.Remove(consumerLogPath); err != nil {
				cLogger.WithError(err).Error("[Verify] fail to delete consumerLogPath")
			} else {
				cLogger.Info("[Verify] delete consumerLogPath")
			}
			if err := fileKit.Remove(producerLogPath); err != nil {
				cLogger.WithError(err).Error("[Verify] fail to delete producerLogPath")
			} else {
				cLogger.Info("[Verify] delete producerLogPath")
			}
		}
	}()

	err = _verify(cLogger, topic, consumerLogPath, producerLogPath, ulid)
	return
}

func _verify(logger *logrus.Logger, topic, consumerLogPath, producerLogPath, ulid string) error {
	ctx0, cancel := context.WithTimeout(context.TODO(), connectTimeout)
	defer cancel()
	consumer, err := NewConsumer(ctx0, pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: fmt.Sprintf("verify_%s_%d", ulid, randomKit.Int(0, 10000000)),
		Type:             pulsar.Exclusive,
	}, consumerLogPath)
	if err != nil {
		return err
	}
	defer consumer.Close()

	ctx1, cancel := context.WithTimeout(context.TODO(), connectTimeout)
	defer cancel()
	producer, err := NewProducer(ctx1, pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	}, producerLogPath)
	if err != nil {
		return err
	}
	defer producer.Close()

	timeStr := timeKit.FormatCurrent(timeKit.FormatCommon)
	texts := []string{
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$0"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$1"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$2"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$3"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$4"),
		fmt.Sprintf("%s&&%s&&%s", ulid, timeStr, "$5"),
	}
	var ch = make(chan struct{}, 1)
	var consumerErrCh = make(chan error, 1)
	var producerErrCh = make(chan error, 1)

	/* consumer */
	consumerCtx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	go func() {
		defer func() {
			logger.Info("[Pulsar, Verify, Consumer] goroutine ends")
		}()

		s := sliceKit.Copy(texts)
		for {
			msg, err := consumer.Receive(consumerCtx)
			if err != nil {
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Pulsar, Verify, Consumer] fail to receive")
				consumerErrCh <- err
				break
			}
			if err := consumer.Ack(msg); err != nil {
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Pulsar, Verify, Consumer] fail to ack")
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
			}).Info("[Pulsar, Verify, Consumer] receive a message")

			if ok && left == 0 {
				logger.Info("[Pulsar, Verify, Consumer] receive all messages!")
				ch <- struct{}{}
				break
			}
		}
	}()

	/* producer */
	go func() {
		defer func() {
			logger.Info("[Pulsar, Verify, Producer] goroutine ends")
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
				logger.WithFields(logrus.Fields{
					"text":  text,
					"error": err.Error(),
				}).Error("[Pulsar, Verify, Producer] fail to send")
				producerErrCh <- err
				break
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[Pulsar, Verify, Producer] succeeded to send")
		}
	}()

	select {
	case <-ch:
		return nil
	case err := <-producerErrCh:
		return err
	case err := <-consumerErrCh:
		return err
	case <-time.After(receiveTimeout):
		return errorKit.New("fail to get all messages within time limit(%s)", receiveTimeout)
	}
}
