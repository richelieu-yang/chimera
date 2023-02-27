package rocketmq5Kit

import "testing"

func TestTestEndpoint(t *testing.T) {
	err := VerifyEndpoint("192.168.80.42:28888", "test")
	if err != nil {
		panic(err)
	}
}
