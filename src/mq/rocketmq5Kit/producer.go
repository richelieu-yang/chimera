package rocketmq5Kit

import (
	"context"
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
)

// NewProducer
/*
@param logConfig 可以为nil
@return rmq_client.Producer实例要手动调用 Start()!!!
*/
func NewProducer(logConfig *LogConfig, config *rmq_client.Config) (rmq_client.Producer, error) {
	lock.Lock()
	defer lock.Unlock()

	if err := logConfig.SetLogout(); err != nil {
		return nil, err
	}

	config, err := processConfig(config)
	if err != nil {
		return nil, err
	}
	return rmq_client.NewProducer(config)
}

// SendMessage
/*
Deprecated: 此方法仅供参考.

@param tag 可以为nil
*/
func SendMessage(producer rmq_client.Producer, topic string, body []byte, tag *string, keys ...string) (*rmq_client.SendReceipt, error) {
	respSlice, err := SendMessageForSendReceipts(producer, topic, body, tag, keys...)
	if err != nil {
		return nil, err
	}
	// SimpleConsumer + Producer的情况下，响应的长度为1
	return respSlice[0], nil
}

// SendMessageForSendReceipts
/*
Deprecated: 此方法仅供参考.
*/
func SendMessageForSendReceipts(producer rmq_client.Producer, topic string, body []byte, tag *string, keys ...string) ([]*rmq_client.SendReceipt, error) {
	msg := &rmq_client.Message{
		Topic: topic,
		Body:  body,
	}
	if tag != nil {
		msg.SetTag(*tag)
	}
	msg.SetKeys(keys...)

	return producer.Send(context.TODO(), msg)
}
