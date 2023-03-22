package pulsarKit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetUp(t *testing.T) {
	pulsarConfig := &Config{
		Addresses: []string{"127.0.0.1:6650"},
		VerifyConfig: &VerifyConfig{
			Topic: "test",
			Print: true,
		},
	}

	err := SetUp(pulsarConfig)
	assert.Nil(t, err)
}
