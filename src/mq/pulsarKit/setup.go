package pulsarKit

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
)

var client pulsar.Client

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		log.Fatal()
	}
}

func SetUp(config *Config) error {
	var err error

	client, err = NewClient(config)
	if err != nil {
		return err
	}

	return nil
}
