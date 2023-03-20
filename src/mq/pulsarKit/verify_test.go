package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyPulsar(t *testing.T) {
	client, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: "pulsar://localhost:6650",
	})
	assert.Nil(t, err)
	assert.NotNil(t, client)

	err = VerifyPulsar(client, "test1", true)
	assert.Nil(t, err)
}
