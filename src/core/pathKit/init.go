package pathKit

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	var err error

	projectDir, err = os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
}
