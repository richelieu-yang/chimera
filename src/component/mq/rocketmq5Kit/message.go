package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
)

// NewMessage
/*
@param tag 可以为nil
@return Producer发送的消息
*/
func NewMessage(topic string, body []byte, tag *string) *rmq_client.Message {
	return &rmq_client.Message{
		Topic: topic,
		Body:  body,
		Tag:   tag,
	}
}
