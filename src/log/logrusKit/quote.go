package logrusKit

import (
	"github.com/sirupsen/logrus"
)

func DisableQuote(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	textFormatter, ok := logger.Formatter.(*logrus.TextFormatter)
	if ok {
		textFormatter.DisableQuote = true
		textFormatter.ForceQuote = false
	}
}

func EnableQuote(logger *logrus.Logger) {
	if logger == nil {
		logger = logrus.StandardLogger()
	}

	textFormatter, ok := logger.Formatter.(*logrus.TextFormatter)
	if ok {
		textFormatter.DisableQuote = false
		textFormatter.ForceQuote = true
	}
}

//// DisableQuoteTemporarily 临时禁用双引号(")
///*
//@param logger 	(1) 可以为nil（此时将采用logrus.StandardLogger()）
//				(2) 只有当formatter为 *logrus.TextFormatter 类型，DisableQuote才会生效
//*/
//func DisableQuoteTemporarily(logger *logrus.Logger, callback func()) {
//	if callback == nil {
//		return
//	}
//
//	if logger == nil {
//		logger = logrus.StandardLogger()
//	}
//	textFormatter, ok := logger.Formatter.(*logrus.TextFormatter)
//	if ok {
//		flag := textFormatter.DisableQuote
//		flag1 := textFormatter.ForceQuote
//
//		textFormatter.DisableQuote = true
//		textFormatter.ForceQuote = false
//
//		callback()
//
//		// restore
//		textFormatter.DisableQuote = flag
//		textFormatter.ForceQuote = flag1
//	} else {
//		callback()
//	}
//}
