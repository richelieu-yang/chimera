package assertKit

import (
	"log"
)

// Must
/*
参考: logx.Must()
*/
func Must(err error) {
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
