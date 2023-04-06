package rocketmq5Kit

import "testing"

func TestVerifyEndpoint(t *testing.T) {
	err := VerifyEndpoint("192.168.80.27:28888;192.168.80.43:28888", "test")
	//err := VerifyEndpoint("127.0.0.1:8081", "test")
	if err != nil {
		panic(err)
	}
}
