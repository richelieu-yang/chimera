package wsKit

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

// setLogger
/*
@param logger1 可以为nil
*/
func setLogger(logger1 *logrus.Logger) {
	logger = logger1
}

// getLogger
/*
@return 必定不为nil
*/
func getLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}
