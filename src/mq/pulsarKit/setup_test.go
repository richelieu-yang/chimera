package pulsarKit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetUp(t *testing.T) {
	pulsarConfig := &Config{
		Addresses: []string{"192.168.80.27:6650", "192.168.80.42:6650", "192.168.80.43:6650"},
		//Addresses: []string{"172.18.21.50:6650"},
		VerifyConfig: &VerifyConfig{
			Topic: "test",
			Print: true,
		},
	}

	err := SetUp(pulsarConfig)
	assert.Nil(t, err)
}
