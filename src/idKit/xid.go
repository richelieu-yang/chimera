package idKit

import (
	"github.com/rs/xid"
	"time"
)

// NewXid
/*
@return 长度(len())为20
*/
func NewXid() string {
	return xid.New().String()
}

func NewWithTime(t time.Time) string {
	return xid.NewWithTime(t).String()
}
