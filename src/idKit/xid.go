package idKit

import (
	"github.com/rs/xid"
	"time"
)

// NewXid
/*
Xid是一个全局唯一的ID生成器，它使用Mongo Object ID算法来生成全局唯一的ID.

PS:
(1) 一个进程（Process）内，生成的id不会重复.

@return 长度(len())固定为20	e.g.  "ckic7hfnl531vbl645n0"
							e.g.1 "ckth51co47mgs2kacmk0"
*/
func NewXid() string {
	return xid.New().String()
}

func NewWithTime(t time.Time) string {
	return xid.NewWithTime(t).String()
}
