package pulsarKit

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/core/timeKit"
	"github.com/richelieu42/chimera/src/idKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/richelieu42/chimera/src/operationKit"
	"github.com/sirupsen/logrus"
	"time"
)

// VerifyPulsar 验证 Pulsar服务 是否启动成功
func VerifyPulsar(client pulsar.Client, topic string, printArgs ...bool) error {
	if client == nil {
		return errorKit.Simple("client == nil")
	}
	if strKit.IsEmpty(topic) {
		return errorKit.Simple("topic is empty")
	}
	if strKit.IsBlank(topic) {
		return errorKit.Simple("topic is blank")
	}

	printFlag := sliceKit.GetFirstItemWithDefault(false, printArgs...)
	logger := logrusKit.NewLogger(nil, operationKit.Ternary(printFlag, logrus.DebugLevel, logrus.PanicLevel))

	var timeLimit = time.Second * 10
	var sendTimeout = time.Second
	//var sendTimeout = time.Nanosecond

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

	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            topic,
		SubscriptionName: ulid,
		Type:             pulsar.Exclusive,
	})
	if err != nil {
		return errorKit.Wrap(err, "fail to create a consumer")
	}
	defer consumer.Close()

	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:       topic,
		SendTimeout: sendTimeout,
	})
	if err != nil {
		return errorKit.Wrap(err, "fail to create a producer")
	}
	defer producer.Close()

	var successCh = make(chan struct{}, 1)
	var errCh0 = make(chan error, 1)
	var errCh1 = make(chan error, 1)

	go func() {
		defer func() {
			logger.Info("[Consumer] goroutine ends")
		}()

		s := sliceKit.Copy(texts)

		for {
			msg, err := consumer.Receive(context.TODO())
			if err != nil {
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Consumer] fail to receive")
				errCh0 <- err
				break
			}
			if err := consumer.Ack(msg); err != nil {
				logger.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Info("[Consumer] fail to ack")
				continue
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
				successCh <- struct{}{}
				break
			}
		}
	}()

	//for cMsg := range consumer.Chan() {
	//	if err := consumer.Ack(cMsg); err != nil {
	//		logger.WithFields(logrus.Fields{
	//			"text": string(cMsg.Payload()),
	//		}).Error("[Consumer] fail to ack")
	//		continue
	//	}
	//	str := string(cMsg.Payload())
	//	var ok bool
	//	if s, ok = sliceKit.Remove(s, str); ok {
	//		left := len(s)
	//		logger.WithFields(logrus.Fields{
	//			"left":  left,
	//			"valid": ok,
	//			"text":  string(cMsg.Payload()),
	//		}).Info("[Consumer] receive a message")
	//		if left == 0 {
	//			successCh <- struct{}{}
	//			break
	//		}
	//	} else {
	//		logger.WithFields(logrus.Fields{
	//			"valid": ok,
	//			"text":  string(cMsg.Payload()),
	//		}).Info("[Consumer] receive a message")
	//	}
	//}

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
				logger.WithFields(logrus.Fields{
					"text":  text,
					"error": err.Error(),
				}).Error("[Producer] fail to send a message")
				errCh1 <- err
				break
			}
			logger.WithFields(logrus.Fields{
				"text": text,
			}).Info("[Producer] send a message")
		}
	}()

	select {
	case <-successCh:
		return nil
	case err := <-errCh1:
		return err
	case <-time.After(timeLimit):
		return errorKit.Simple("fail to get all messages within time limit(%s)", timeLimit)
	}
}
