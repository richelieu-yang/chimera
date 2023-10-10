package idKit

import (
	"github.com/rs/xid"
	"time"
)

// NewXid
/*
@return 长度(len())固定为20.	e.g."ckic7hfnl531vbl645n0"
*/
func NewXid() string {
	return xid.New().String()
}

func NewWithTime(t time.Time) string {
	return xid.NewWithTime(t).String()
}
