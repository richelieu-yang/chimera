package rocketmq5Kit

import "testing"

func TestTestEndpoint(t *testing.T) {
	err := VerifyEndpoint("localhost:8081", "test")
	if err != nil {
		panic(err)
	}
}
