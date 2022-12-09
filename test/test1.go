package main

import rmq_client "github.com/apache/rocketmq-clients/golang"

func main() {
	msg := &rmq_client.Message{}
	msg.SetTag("")
	//msg.SetKeys(keys...)
}
