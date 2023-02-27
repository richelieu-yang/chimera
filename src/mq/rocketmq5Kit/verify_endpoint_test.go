package rocketmq5Kit

import "testing"

func TestVerifyEndpoint(t *testing.T) {
	err := VerifyEndpoint("127.0.0.1:8081", "test")
	if err != nil {
		panic(err)
	}
}
