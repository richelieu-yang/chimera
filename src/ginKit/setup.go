package ginKit

import (
	"github.com/sirupsen/logrus"
)

func MustSetUp(config *Config) {
	err := SetUp(config)
	if err != nil {
		logrus.Fatal(err)
	}
}

func SetUp(config *Config) error {

	// TODO:
	return nil
}
