package mysqlKit

import "github.com/sirupsen/logrus"

func MustSetUp(config *Config) {
	if err := SetUp(config); err != nil {
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *Config) error {

	return nil
}
