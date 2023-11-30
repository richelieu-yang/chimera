package carbonKit

import "github.com/golang-module/carbon/v2"

var (
	Parse func(value string, timezone ...string) carbon.Carbon = carbon.Parse

	ParseByLayout func(value, layout string, timezone ...string) carbon.Carbon = carbon.ParseByLayout

	ParseByFormat func(value, format string, timezone ...string) carbon.Carbon = carbon.ParseByFormat
)
