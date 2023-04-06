package timeKit

import (
	"github.com/dustin/go-humanize"
	"time"
)

// ToRelativeString
/*
e.g.
timeKit.ToRelativeString(time.Now().Add(time.Hour * -16))					=> "16 hours ago"
timeKit.ToRelativeString(time.Now().Add(time.Hour * 16)) 					=> "15 hours from now"
timeKit.ToRelativeString(time.Now().Add(time.Hour * 24 * 21))				=> "2 weeks from now"
timeKit.ToRelativeString(time.Now().Add(time.Hour*24*21 + time.Second))		=> "3 weeks from now"
*/
func ToRelativeString(t time.Time) string {
	return humanize.Time(t)
}
